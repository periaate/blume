package blume

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/pred/has"
)

func Args(preds ...func([]string) bool) Option[Array[String]] {
	var res []string
	if len(os.Args) >= 1 {
		res = os.Args[1:]
	}
	ok := PredAnd(preds...)(res)
	if !ok {
		return None[Array[String]]()
	}
	return Some(ToArray(Map(func(s string) String { return String(s) })(res)))
}

func Piped(f *os.File, preds ...func([]string) bool) Option[Array[String]] {
	ok := has.Pipe(f)
	if !ok {
		return None[Array[String]]()
	}
	var res []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	ok = PredAnd(preds...)(res)
	if !ok {
		return None[Array[String]]()
	}
	return Some(ToArray(Map(func(s string) String { return String(s) })(res)))
}

func Stringify(s string) String { return String(s) }

func Lines(bar []byte) []string {
	scanner := bufio.NewScanner(Buf(bar))
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

type Directory string

func Dir[S ~string](root S) Directory { return Directory(root) }

func (d Directory) Read() Result[Array[String]] {
	if res, err := fsio.ReadDir(string(d)); err == nil {
		return Ok(ToArray(Map(func(file fsio.Entry) String {
			return String(file.Path())
		})(res)))
	} else {
		return Err[Array[String]](err.Error())
	}
}

func (d Directory) First(pred Pred[string]) Option[String] {
	if res, ok := fsio.First(string(d), pred); ok {
		return Some(String(res))
	}
	return None[String]()
}

func (d Directory) Find(pred Pred[string]) Option[Array[String]] {
	if res := fsio.Find(string(d), pred); len(res) > 0 {
		return Some(ToArray(Map(Stringify)(res)))
	}
	return None[Array[String]]()
}

func (d Directory) Ascend(pred Pred[string]) Option[String] {
	if res, ok := fsio.Ascend(string(d), pred); ok {
		return Some(String(res))
	}
	return None[String]()
}

func Read[S ~string](sar ...S) Result[String] {
	str := Path(sar...)
	bar, err := os.ReadFile(string(str))
	if err != nil {
		return Err[String](err.Error())
	}
	return Ok(String(bar))
}

func Path[S ~string](sar ...S) S {
	return S(filepath.Join(ToArray(Map(StoD[S])(sar)).Map(ReplacePrefix("~", Must(os.UserHomeDir()))).Value...))
}
