package fsio

import (
	"os"
	"path/filepath"
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

// Entry is an os.DirEntry wrapper which is aware of it's full path, not just its name.
type Entry struct {
	os.DirEntry
	root string
}

func (e Entry) Path() string { return filepath.Join(e.root, e.Name()) }

func ReadDir(root string) (res []Entry, err error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return res, err
	}
	res = make([]Entry, 0, len(entries))
	for _, v := range entries {
		res = append(res, Entry{v, root})
	}
	return res, nil
}

func Traverse(root string, walk func(path Entry) (skip, stop bool)) (err error) {
	root, err = filepath.Abs(root)
	if err != nil {
		return
	}
	queue := []string{root}
	var item string
	for len(queue) > 0 {
		item, queue = queue[0], queue[1:]
		entries, err := ReadDir(item)
		if err != nil {
			return err
		}
		for _, e := range entries {
			skip, stop := walk(e)
			if stop {
				return nil
			}
			if skip {
				continue
			}
			if IsDir(e.Path()) {
				queue = append(queue, e.Path())
			}
		}
	}
	return nil
}

func First(root string, cond func(string) bool) (res string, ok bool) {
	Traverse(root, func(entry Entry) (skip bool, stop bool) {
		res = entry.Path()
		ok = cond(res)
		return false, ok
	})
	return
}

func Find(root string, cond func(string) bool) (res []string) {
	Traverse(root, func(entry Entry) (_ bool, _ bool) {
		path := entry.Path()
		if cond(path) {
			res = append(res, path)
		}
		return
	})
	return
}

func Ascend(path string, cond func(string) bool) (res string, ok bool) {
	for {
		if path == filepath.Dir(path) {
			return
		}
		entries, err := ReadDir(path)
		if err != nil {
			return
		}
		for _, entry := range entries {
			res = entry.Path()
			ok = cond(res)
			if ok {
				return
			}
		}
		path = filepath.Dir(path)
	}
}
