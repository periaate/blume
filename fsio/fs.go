package fsio

import (
	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/yap"
)

func FindFirst[A, S ~string](root A, preds ...Predicate[S]) Option[FilePath] {
	type queueItem struct {
		path FilePath
	}

	pred := PredAnd(preds...)

	queue := []queueItem{{path: FilePath(root)}}
	visited := make(map[string]bool)
	visited[string(root)] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		dirRes := ReadsDir(item.path)
		if dirRes.IsErr() { continue }

		entries := dirRes.Unwrap()
		for _, e := range entries.Values() {
			eStr := e.String()
			if !visited[eStr] {
				visited[eStr] = true
				if e.HasSuffix("/") { queue = append(queue, queueItem{path: e}) }
				if pred(S(e)) { return Some(e) }
			}
		}
	}

	return None[FilePath]()
}

func Find[A, S ~string](root A, preds ...Predicate[S]) Option[Array[FilePath]] {
	type queueItem struct {
		path FilePath
	}
	pred := PredAnd(preds...)

	res := Arr[FilePath]{}

	queue := []queueItem{{path: FilePath(root)}}
	visited := make(map[string]bool)
	visited[string(root)] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		dirRes := ReadsDir(item.path)
		if dirRes.IsErr() { continue }

		entries := dirRes.Unwrap()
		for _, e := range entries.Values() {
			eStr := e.String()
			if !visited[eStr] {
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
	}

	if res.Len() == 0 { None[Array[FilePath]]() }
	return Some[Array[FilePath]](res)
}

func Ascend[A, S ~string](root A, preds ...Predicate[S]) Option[FilePath] {
	fp := FilePath(root)
	pred := PredAnd(preds...)
	for {
		if fp == fp.Dir() { return None[FilePath]() }
		last := fp
		res, err := ReadsDir(fp).Values()
		if err != nil { return None[FilePath](err) } 
		rfp, err := res.First(func(f FilePath) bool { return pred(S(f)) }).Values()
		if err == nil { return Some(rfp) }
		fp = fp.Dir()
		yap.Debug("ascending", "from", last, "to", fp)
	}
}
