package fsio

import (
	. "github.com/periaate/blume"
)

func FindFirst[A, S ~string](root A, preds ...FnA[S, bool]) Option[String] {
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

		dirRes := ReadDir(item.path)
		if !dirRes.Ok() { continue }

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

	return None[String]()
}

func Find[A, S ~string](root A, preds ...FnA[S, bool]) Option[Array[String]] {
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

		dirRes := ReadDir(item.path)
		if !dirRes.Ok() { continue }

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

	if res.Len() == 0 { None[Array[String]]() }
	return Some[Array[String]](res)
}

func Ascend[A, S ~string](root A, preds ...FnA[S, bool]) Option[S] {
	fp := String(root)
	pred := PredAnd(preds...)
	for {
		if fp == Dir(fp) { return None[S]() }
		tn := ReadDir(S(fp))
		if !tn.Ok() { return None[S]() }
		opt := tn.Unwrap().First(pred)
		if opt.Ok() { return Some(opt.Unwrap()) }
		fp = Dir(fp)
	}
}
