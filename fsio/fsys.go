package fsio

import (
	"os"
	"strings"

	. "github.com/periaate/blume"
)

// ReadDir reads the directory and returns a list of files.
func ReadDir[S ~string](inp S) (res Array[S], err error) {
	f := string(inp)
	if HasPrefix("~")(f) { f = strings.Replace(f, "~", Home(), 1) }
	if !IsDir(f) { return Err[Array[S]]("{:s} is not a directory", f) }
	
	entries, err := os.ReadDir(f)
	if err != nil { return Err[Array[S]]("failed to read directory [{:s}] with error: [{:w}]", f, err) }

	arr := make([]S, 0, len(entries))
	for _, entry := range entries {
		fp := entry.Name()
		if entry.IsDir() { fp += "/" }
		joined := Join(f, fp)
		if len(joined) == 0 { return Err[Array[S]]("path {:s} with {:s} has length of 0", f, fp) }
		arr = append(arr, S(joined))
	}

	return Ok(ToArray(arr))
}



func FindFirst[A, S ~string](root A, preds ...func(S) bool) Option[String] {
	type queueItem struct {
		path String
	}

	pred := PredAnd(preds...)

	queue := []queueItem{{path: String(root)}}
	visited := make(map[string]bool)
	visited[string(root)] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		entries, err := ReadDir(item.path)
		if err != nil { continue }
		for _, e := range entries.Values() {
			eStr := e.String()
			if !visited[eStr] {
				visited[eStr] = true
				if e.HasSuffix("/") { queue = append(queue, queueItem{path: e}) }
				if pred(S(e)) { return Some(e) }
			}
		}
	}

	return None[String]()
}

func Find[A, S ~string](root A, preds ...func(S) bool) Option[Array[String]] {
	type queueItem struct {
		path String
	}
	pred := PredAnd(preds...)

	res := Array[String]{}

	queue := []queueItem{{path: String(root)}}
	visited := make(map[string]bool)
	visited[string(root)] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		entries, err := ReadDir(item.path)
		if err != nil { continue }
		for _, e := range entries.Values() {
			eStr := e.String()
			if visited[eStr] { continue }
			visited[eStr] = true
			b := true
			if !pred(S(e)) {
				b = false
				break
			}
			if b {
				if e.HasSuffix("/") { queue = append(queue, queueItem{path: e}) }
				res = res.Append(e)
			}
		}
	}

	if res.Len() == 0 { return None[Array[String]]() }
	return Some(res)
}

func Ascend[A, S ~string](root A, preds ...func(S) bool) Option[S] {
	fp := String(root)
	pred := PredAnd(preds...)
	for {
		if fp == Dir(fp) { return None[S]() }
		tn, err := ReadDir(S(fp))
		if err != nil { return None[S]() }
		if opt := tn.First(pred); opt.Ok { return opt }
		fp = Dir(fp)
	}
}

