package str

import (
	"fmt"
	"testing"
)

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
		if got, _ := ReplacePrefix(c.inp, c.pats...); got != c.exp {
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
		if got, _ := ReplaceSuffix(c.inp, c.pats...); got != c.exp {
			t.Fatalf("expected %s, got %s", c.exp, got)
		}
	}
}

func TestSplitWithAll(t *testing.T) {
	tst := `(foo (bar baz abc))`
	delims := []string{"(", ")", " "}
	res := Split(tst, true, delims...)
	for i, r := range res {
		fmt.Println(i+1, r)
	}
	if len(res) != 11 {
		t.Fatalf("expected 6, got %d", len(res))
	}
}
