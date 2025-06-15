package blume

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/pred/has"
)

func Input[S ~string](from ...string) Array[string] {
	res := []string{}
	for _, arg := range from {
		switch strings.ToLower(arg) {
		case "args":
			res = append(res, Args()...)
		case "pipe", "piped":
			res = append(res, Piped(os.Stdin).Must()...)
		}
	}

	return res
}

func AllArgs(n ...int) Array[string] { return JoinAfter(Args())(Piped(os.Stdin).OrDef()) }

func Args(n ...int) []string {
	var res []string
	if len(os.Args) >= 1 {
		res = os.Args[1:]
	}
	if len(n) == 0 {
		return res
	}
	if len(res) < n[0] {
		return res
	}
		
	return res[n[0]:]
}

func Arg(n int) Option[string] {
	if len(os.Args) > n+1 {
		return Some(S(os.Args[n+1]))
	}

	return None[string]()
}

func Piped(input ...*os.File) Option[Array[string]] {
	var f *os.File
	if len(input) == 0 { f = os.Stdin } else { f = input[0] }

	if !has.Pipe(f) { return None[Array[string]]() }
	return Some(Lines(f))
}

func stringify(s string) string { return string(s) }

func Lines[B any](bar B) A[S] {
	scanner := bufio.NewScanner(Buf(any(bar)))
	res := []string{}
	for scanner.Scan() {
		res = append(res, S(scanner.Text()))
	}
	return res
}

func Entries(s S) Result[Array[string]] {
	if res, err := fsio.ReadDir(s); err == nil {
		return Ok(Map[fsio.Entry, S](func(file fsio.Entry) string {
			return string(file.Path())
		})(res))
	} else {
		return Err[Array[string]](err.Error())
	}
}
//
// func (d string) FirstFS(pred Pred[string]) Option[string] {
// 	if res, ok := fsio.First(string(d), func(s string) bool { return pred(S(s))}); ok {
// 		return Some(string(res))
// 	}
// 	return None[string]()
// }
//
// func (d S) FindFS(pred Pred[string]) Option[Array[string]] {
// 	if res := fsio.Find(string(d), pred); len(res) > 0 {
// 		return Some(Map[string, S](stringify)(res))
// 	}
// 	return None[Array[string]]()
// }

// func (d string) AscendFS(pred Pred[string]) Option[string] {
// 	if res, ok := fsio.Ascend(string(d), pred); ok {
// 		return Some(string(res))
// 	}
// 	return None[string]()
// }

func Read(sar ...S) Result[string] {
	str := Path(sar...)
	bar, err := os.ReadFile(string(str))
	if err != nil {
		return Err[string](err.Error())
	}
	return Ok(string(bar))
}

func Reads(filepath string) string { return Read(filepath).Must() }

// func (s S) And(fns ...func(S) bool) bool { return PredAnd(fns...)(s) }

func Path(sar ...S) string {
	var fp S
	sar = Map[S, S](Replace("~", S(Must(os.UserHomeDir()))))(sar)
	fps := filepath.Join(Into[[]string](sar).Value...)
	absFp, err := filepath.Abs(fps)
	if err != nil { fp = S(fps) } else { fp = S(absFp) }
	if IsDir(fp) { fp = EnsureSuffix("/")(fp) }
	return fp
}

// LPath resolves symlinks
func LPath(sar ...S) string {
	sar = Map[S, S](Replace("~", S(Must(os.UserHomeDir()))))(sar)
	fp := filepath.Join(From[[]S, []string](sar).Value...)
	if IsSymlink(S(fp)).Value {
		evaluated, err := filepath.EvalSymlinks(fp)
		if err != nil {
			return S(fp)
		}
		fp = evaluated
	}
	absFp, err := filepath.Abs(fp)
	if err == nil {
		return absFp
	}
	return fp
}

func Paths(v string, sar ...string) string {
	sar = Map[S, S](Replace("~", S(Must(os.UserHomeDir()))))(sar)
	fp := filepath.Join(From[[]S, []string](sar).Value...)
	absFp, err := filepath.Abs(fp)
	if err == nil {
		return string(absFp)
	}
	return string(fp)
}

func AppendTo(path S) (res Result[*os.File]) {
	return res.Auto(os.OpenFile(Path(path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644))
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
