package str

import (
	"testing"
)

func TestContains(t *testing.T) {
	expect := []struct {
		tar  string
		args []string
	}{
		{"Hello, World", []string{"llo ", "he", ", W"}},
	}

	for _, ex := range expect {
		if !Contains(ex.args...)(ex.tar) {
			t.Fatal("match not found", ex.tar, ex.args)
		}
	}
}

func TestPre(t *testing.T) {
	expect := []struct {
		tar  string
		args []string
	}{
		{"Hello, World", []string{"llo ", "he", ", W", "He"}},
	}

	for _, ex := range expect {
		if !HasPrefix(ex.args...)(ex.tar) {
			t.Fatal("match not found", ex.tar, ex.args)
		}
	}
}

func TestSuf(t *testing.T) {
	expect := []struct {
		tar  string
		args []string
	}{
		{"Hello, World", []string{"llo ", "he", ", W", "ld"}},
	}

	for _, ex := range expect {
		if !HasSuffix(ex.args...)(ex.tar) {
			t.Fatal("match not found", ex.tar, ex.args)
		}
	}
}

func TestReplacePre(t *testing.T) {
	cases := []struct {
		inp  string
		pats []string

		exp string
	}{
		{"./abc/./hello.jpg", []string{"./", "hii/", "abc", "no"}, "hii/abc/./hello.jpg"},
		{"./abc/./hello.jpg", []string{"/", "no", "./abc", "yay!"}, "yay!/./hello.jpg"},
		{".", []string{"./", "./", `.\`, "./", ".", "./"}, "./"},
		{"", []string{"./", "./", `.\`, "./", ".", "./"}, ""},
		{"", []string{}, ""},
		{"abc", []string{}, "abc"},
	}

	for _, c := range cases {
		if got := ReplacePrefix(c.pats...)(c.inp); got != c.exp {
			t.Fatalf("expected %s, got %s", c.exp, got)
		}
	}
}

func TestReplaceSuf(t *testing.T) {
	cases := []struct {
		inp  string
		pats []string

		exp string
	}{
		{"./abc/./hello.jpg", []string{".jpg", ".png", "jpg", "no"}, "./abc/./hello.png"},
		{"./abc/./hello.jpg", []string{".jpeg", "no", "hello.jpg", "world!!"}, "./abc/./world!!"},
	}

	for _, c := range cases {
		if got := ReplaceSuffix(c.pats...)(c.inp); got != c.exp {
			t.Fatalf("expected %s, got %s", c.exp, got)
		}
	}
}
