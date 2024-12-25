package fsio

import (
	"os"
	"path/filepath"

	"github.com/periaate/blume/pred"
)

// Name returns the file name without the extension and directory.
func Name(f string) string {
	b := filepath.Base(f)
	r := b[:len(b)-len(filepath.Ext(b))]
	return filepath.Clean(r)
}

// IsDir checks if input is a directory.
func IsDir(f string) bool {
	info, err := os.Stat(f)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Exists checks if the input exists.
func Exists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

type Entry struct {
	os.DirEntry
	root string
}

func (e Entry) Path() string { return filepath.Join(e.root, e.Name()) }

// ReadDir reads the directory and returns a list of files.
func ReadDir(root string) (res []Entry, err error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return
	}

	res = make([]Entry, 0, len(entries))
	for _, v := range entries {
		res = append(res, Entry{v, root})
	}

	return
}

func Traverse(root string, walk func(path Entry) (skip, stop bool)) {
	queue := []string{root}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		entries, err := ReadDir(item)
		if err != nil {
			continue
		}
		for _, e := range entries {
			skip, stop := walk(e)
			if stop {
				return
			}
			if skip {
				continue
			}
			path := e.Path()
			if IsDir(path) {
				queue = append(queue, path)
			}
		}
	}
}

func First(root string, pred func(string) bool) (res string, ok bool) {
	Traverse(root, func(entry Entry) (skip bool, stop bool) {
		path := entry.Path()
		if pred(path) {
			res = path
			ok = true
			return false, true
		}
		return
	})
	return
}

func Find(root string, pred func(string) bool) (res []string) {
	Traverse(root, func(entry Entry) (_ bool, _ bool) {
		path := entry.Path()
		if pred(path) {
			res = append(res, path)
		}
		return
	})
	return
}

func Ascend(root string, preds ...func(string) bool) (res string, ok bool) {
	fp := root
	pred := pred.And(preds...)
	for {
		if fp == filepath.Dir(fp) {
			return
		}
		tn, err := ReadDir(fp)
		if err != nil {
			return
		}
		for _, entry := range tn {
			path := entry.Path()
			if pred(path) {
				return path, true
			}
		}
		fp = filepath.Dir(fp)
	}
}
