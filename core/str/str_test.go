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
