package blume

import (
	"os"
	"regexp"
	"time"

	"github.com/farmergreg/rfsnotify"
	"github.com/fsnotify/fsnotify"
)

func Auto[A any](value A, handles ...any) Result[A] {
	if IsOk(value, handles...) {
		return Ok(value)
	}
	return Err[A](handles...)
}

func (s String) Stat() Result[os.FileInfo] { return Auto(os.Stat(s.String())) }
func (s String) ModTime() Result[time.Time] {
	r := s.Stat()
	if !r.IsOk() {
		return Err[time.Time](r.Other)
	}
	return Ok(r.Value.ModTime())
}

func (s String) ModMilli() Result[int64] {
	r := s.Stat()
	if !r.IsOk() {
		return Err[int64](r.Other)
	}
	return Ok(r.Value.ModTime().UnixMilli())
}

func (p String) MkdirAll(n os.FileMode) Result[String] {
	if p.Path().Exists() && p.Path().IsDir() {
		return Ok(p)
	}
	return Auto(p, os.MkdirAll(p.String(), n))
}
func (p String) WriteFile(bytes []byte, n os.FileMode) Result[String] {
	return Auto(p, os.WriteFile(p.String(), bytes, n))
}

func SplitRegex(pattern String) func(input String) []String {
	return func(input String) []String {
		re := regexp.MustCompile(pattern.String())
		indexes := re.FindAllStringIndex(input.String(), -1)
		if len(indexes) == 0 {
			return []String{input}
		}

		result := make([]String, 0, 2*len(indexes)+1)
		lastEnd := 0

		for _, idx := range indexes {
			start, end := idx[0], idx[1]
			if start > lastEnd {
				result = append(result, input[lastEnd:start])
			}
			result = append(result, input[start:end])
			lastEnd = end
		}

		if lastEnd < len(input) {
			result = append(result, input[lastEnd:])
		}

		return result
	}
}

func (s S) Listen(fn func(s S), recursive bool, ops ...fsnotify.Op) S {
	Listen(fn, recursive, ops...)(s)
	return s
}

func Listen(fn func(s S), recursive bool, ops ...fsnotify.Op) func(S) {
	rw := Auto(rfsnotify.NewWatcher()).Must()
	f := Is(ops...)
	go func() {
		for {
			ev := <-rw.Events
			if f(fsnotify.Op(ev.Op)) {
				fn(S(ev.Name))
			}
		}
	}()

	return func(s S) {
		if recursive {
			Auto(rw.AddRecursive(s.String())).Must()
		} else {
			Auto(rw.Add(s.String())).Must()
		}
	}
}
