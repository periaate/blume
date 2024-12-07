package fsio

import (
	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
	"github.com/periaate/blume/yap"
)

// Find performs a BFS directory search starting from "root",
// returning the first file that satisfies one of the provided predicates.
func Find[A, S ~string](root A, preds ...T.Predicate[S]) T.Result[FilePath] {
	type queueItem struct {
		path FilePath
	}

	pred := gen.PredAnd(preds...)

	queue := []queueItem{{path: FilePath(root)}}
	visited := make(map[string]bool)
	visited[string(root)] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		dirRes := ReadsDir(item.path)
		if dirRes.IsErr() {
			continue
		}

		entries := dirRes.Unwrap()
		for _, e := range entries.Array() {
			eStr := e.String()
			if !visited[eStr] {
				visited[eStr] = true
				if e.HasSuffix("/") {
					queue = append(queue, queueItem{path: e})
				}
				if pred(S(e)) {
					return T.Results(e, nil)
				}
			}
		}
	}

	return T.Results(FilePath(""), "no matches found")
}

func Filter[A, S ~string](root A, preds ...T.Predicate[S]) T.Result[gen.Array[FilePath]] {
	type queueItem struct {
		path FilePath
	}
	pred := gen.PredAnd(preds...)

	res := gen.Array[FilePath]{}

	queue := []queueItem{{path: FilePath(root)}}
	visited := make(map[string]bool)
	visited[string(root)] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		dirRes := ReadsDir(item.path)
		if dirRes.IsErr() {
			continue
		}

		entries := dirRes.Unwrap()
		for _, e := range entries.Array() {
			eStr := e.String()
			if !visited[eStr] {
				visited[eStr] = true
				b := true
				if !pred(S(e)) {
					b = false
					break
				}
				if b {
					if e.HasSuffix("/") {
						queue = append(queue, queueItem{path: e})
					}

					res = res.Append(e)
				}
			}
		}
	}

	if res.Len() == 0 {
		return T.Results(res, "no matches found")
	}
	return T.Results(res, nil)
}

func Ascend[A, S ~string](root A, preds ...T.Predicate[S]) T.Result[FilePath] {
	fp := FilePath(root)
	pred := gen.PredAnd(preds...)
	for {
		if fp == fp.Dir() {
			return T.Results(FilePath(""), "couldn't ind a  match while ascending")
		}
		last := fp
		res, err := ReadsDir(fp).Values()
		if err != nil {
			return T.Results(FilePath(""), err)
		}
		rfp, err := res.First(func(f FilePath) bool { return pred(S(f)) }).Values()
		if err == nil {
			return T.Results(rfp, nil)
		}
		fp = fp.Dir()
		yap.Debug("ascending", "from", last, "to", fp)
	}
}
