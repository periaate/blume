package blume

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/pred/has"
)

func Input[S ~string](from ...string) Array[String] {
	res := []String{}
	for _, arg := range from {
		switch String(arg).ToLower() {
		case "args":
			res = append(res, Args().Value...)
		case "pipe", "piped":
			res = append(res, Piped(os.Stdin).Must().Value...)
		}
	}

	return ToArray(res)
}

func AllArgs(n ...int) Array[String] { return Args(n...).JoinAfter(Piped(os.Stdin).OrDef()) }

func Args(n ...int) Array[String] {
	var res []string
	if len(os.Args) >= 1 {
		res = os.Args[1:]
	}
	if len(n) == 0 {
		return DAS(res...)
	}
	if len(res) < n[0] {
		return Arr[String]()
	}
	return DAS(res[n[0]:]...)
}

func Arg(n int) Option[String] {
	if len(os.Args) > n+1 {
		return Some(S(os.Args[n+1]))
	}

	return None[String]()
}

func Piped(input ...*os.File) Option[Array[String]] {
	var f *os.File
	if len(input) == 0 { f = os.Stdin } else { f = input[0] }

	if !has.Pipe(f) { return None[Array[String]]() }
	return Some(ToArray(Lines(f)))
}

func Stringify(s string) String { return String(s) }

func Lines[B any](bar B) []String {
	scanner := bufio.NewScanner(Buf(any(bar)))
	res := []String{}
	for scanner.Scan() {
		res = append(res, S(scanner.Text()))
	}
	return res
}

func (d String) ServeFS() http.Handler {
	s := S(d)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Path(S(r.URL.EscapedPath())).Rep(Rgx(`(\.+)`), ".", Pre("~"), "").Open().Then(func(f *os.File) *os.File {
			defer f.Close()
			io.Copy(w, f)
			return f
		})
	})
}

func Muxes(src, pat S, muxs ...*http.ServeMux) *http.ServeMux {
	var mux *http.ServeMux
	if len(muxs) > 0 {
		mux = muxs[0]
	}
	if mux == nil { mux = http.DefaultServeMux }
	rgx1 := ReplaceRegex(`(\.+)`, ".")
	rgx2 := ReplaceRegex("~", "")
	rep := ReplacePrefix(pat, "")
	s := string(src)
	mux.Handle(string(pat.Println()), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := rgx1(S(r.URL.EscapedPath()))
		path = rgx2(path)
		path = rep(path)

		f, err := os.Open(filepath.Join(s, string(path)))
		if err != nil {
			f, err = os.Open(filepath.Join(s, string(path.EnsureSuffix("/")) + "index.html"))
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		}

		defer f.Close()
		io.Copy(w, f) 
	}))
	return mux
}

func Entries(s S) Result[Array[String]] { return s.Entries() }
func (d String) Entries() Result[Array[String]] {
	if res, err := fsio.ReadDir(string(d)); err == nil {
		return Ok(ToArray(Map(func(file fsio.Entry) String {
			return String(file.Path())
		})(res)))
	} else {
		return Err[Array[String]](err.Error())
	}
}

func (d String) FirstFS(pred Pred[String]) Option[String] {
	if res, ok := fsio.First(string(d), func(s string) bool { return pred(S(s))}); ok {
		return Some(String(res))
	}
	return None[String]()
}

func (d S) FindFS(pred Pred[string]) Option[Array[String]] {
	if res := fsio.Find(string(d), pred); len(res) > 0 {
		return Some(ToArray(Map(Stringify)(res)))
	}
	return None[Array[String]]()
}

func (d String) AscendFS(pred Pred[string]) Option[String] {
	if res, ok := fsio.Ascend(string(d), pred); ok {
		return Some(String(res))
	}
	return None[String]()
}

func Read(sar ...S) Result[String] {
	str := Path(sar...)
	bar, err := os.ReadFile(string(str))
	if err != nil {
		return Err[String](err.Error())
	}
	return Ok(String(bar))
}

func Reads(filepath String) String { return Read(filepath).Must() }

func (s S) And(fns ...func(S) bool) bool { return PredAnd(fns...)(s) }

func Path(sar ...S) String {
	sar = Map(Replace("~", S(Must(os.UserHomeDir()))))(sar)
	fps := filepath.Join(SS[S, string](sar)...)
	absFp, _ := filepath.Abs(fps)
	fp := S(T(!IsZero(absFp), absFp, fps))
	if fp.IsDir() { fp.EnsureSuffix("/") }
	return fp
}

// LPath resolves symlinks
func LPath(sar ...S) String {
	sar = Map(Replace("~", S(Must(os.UserHomeDir()))))(sar)
	fp := filepath.Join(SS[S, string](sar)...)
	if IsSymlink(S(fp)).Value {
		evaluated, err := filepath.EvalSymlinks(fp)
		if err != nil {
			return S(fp)
		}
		fp = evaluated
	}
	absFp, err := filepath.Abs(fp)
	if err == nil {
		return String(absFp)
	}
	return String(fp)
}

func Paths(v String, sar ...String) String {
	sar = Map(Replace("~", S(Must(os.UserHomeDir()))))(sar)
	fp := filepath.Join(SS[S, string](sar)...)
	absFp, err := filepath.Abs(fp)
	if err == nil {
		return String(absFp)
	}
	return String(fp)
}

func R[A any](val A, err error) Result[A] {
	if IsOk(val, err) {
		return Ok(val)
	} else {
		return Err[A](err.Error())
	}
}

func AppendTo[S ~string](path S) Result[*os.File] {
	return R(os.OpenFile(string(path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644))
}
func AppendLog[A any](f *os.File) (*os.File, func(a A) A) {
	mut := sync.Mutex{}
	return f,
		func(a A) A {
			mut.Lock()
			defer mut.Unlock()
			f.Write([]byte(fmt.Sprintf("%v\n", a)))
			return a
		}
}
