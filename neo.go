package blume

import (
	// "net/http"
	// "os"
	"regexp"
	// "sync"
	// "time"

	// "github.com/farmergreg/rfsnotify"
	// "github.com/fsnotify/fsnotify"
)

// func (s String) Stat() (res Result[os.FileInfo]) { return res.Auto(os.Stat(s.String())) }
// func (s String) ModTime() (res Result[time.Time]) {
// 	r := s.Stat()
// 	if !r.IsOk() {
// 		return Err[time.Time](r.Other)
// 	}
// 	return Ok(r.Value.ModTime())
// }
//
// func (s String) ModMilli() Result[int64] {
// 	r := s.Stat()
// 	if !r.IsOk() {
// 		return Err[int64](r.Other)
// 	}
// 	return Ok(r.Value.ModTime().UnixMilli())
// }
//
// func (p String) MkdirAll(n os.FileMode) (res Result[String]) {
// 	if p.Path().Exists() && p.Path().IsDir() {
// 		return Ok(p)
// 	}
// 	return res.Auto(p, os.MkdirAll(p.String(), n))
// }
//
// func (p String) WriteFile(bytes []byte, n os.FileMode) (res Result[String]) {
// 	return res.Auto(p, os.WriteFile(p.String(), bytes, n))
// }

// SplitRegex keeps matches
func SplitRegex(pattern string) func(input string) []string {
	return func(input string) []string {
		re := regexp.MustCompile(pattern)
		indexes := re.FindAllStringIndex(input, -1)
		if len(indexes) == 0 {
			return []string{input}
		}

		result := make([]string, 0, 2*len(indexes)+1)
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

// func (s S) Listen(fn func(s S), recursive bool, ops ...fsnotify.Op) S {
// 	Listen(fn, recursive, ops...)(s)
// 	return s
// }
//
// func Listen(fn func(s S), recursive bool, ops ...fsnotify.Op) func(S) {
// 	rw := R(rfsnotify.NewWatcher()).Must()
// 	f := Is(ops...)
// 	go func() {
// 		for {
// 			ev := <-rw.Events
// 			// I don't know why, but many create events are suffixed with `~`, as well as many events being duplicated
// 			if f(fsnotify.Op(ev.Op)) && !HasSuffix("~")(S(ev.Name)) {
// 				fn(S(ev.Name))
// 			}
// 		}
// 	}()
//
// 	return func(s S) {
// 		s = s.Path()
// 		if recursive {
// 			err := rw.AddRecursive(s.String())
// 			if err != nil {
// 				panic("AddRecursive return non nil error: " + err.Error())
// 			}
// 		} else {
// 			R(rw.Add(s.String())).Must()
// 		}
// 	}
// }

// func DebounceMap[K comparable](callback func(K), dur time.Duration) func(K) {
// 	var mu sync.Mutex
// 	pending := make(map[K]time.Time)
//
// 	go func() {
// 		ticker := time.NewTicker(dur / 2)
// 		defer ticker.Stop()
//
// 		for range ticker.C {
// 			now := time.Now()
// 			mu.Lock()
// 			for k, timestamp := range pending {
// 				if now.Sub(timestamp) >= dur {
// 					go callback(k)
// 					delete(pending, k)
// 				}
// 			}
// 			mu.Unlock()
// 		}
// 	}()
//
// 	return func(key K) {
// 		mu.Lock()
// 		pending[key] = time.Now()
// 		mu.Unlock()
// 	}
// }
