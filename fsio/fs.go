package fsio

import (
	. "github.com/periaate/blume"
)

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

	if res.Len() == 0 { None[Array[String]]() }
	return Some[Array[String]](res)
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
