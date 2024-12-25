package has

import "testing"

func TestAny(t *testing.T) {
	expect := []struct {
		tar  string
		args []string
	}{
		{"Hello, World", []string{"llo ", "he", ", W"}},
	}

	for _, ex := range expect {
		if !Any(ex.args...)(ex.tar) {
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
		if !Prefix(ex.args...)(ex.tar) {
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
		if !Suffix(ex.args...)(ex.tar) {
			t.Fatal("match not found", ex.tar, ex.args)
		}
	}
}
