package gen

import "testing"

func TestSplit(t *testing.T) {
	fn := func(s string) bool { return s == "::" }
	args := []string{"a", "b", "::", "c", "d", "::", "e", "f"}
	want := [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	got := Split(fn)(args)
	for i, v := range got {
		if len(v) != len(want[i]) {
			t.Errorf("Split(%v) = %v; want %v", args, got, want)
		}
		for j, w := range v {
			if w != want[i][j] {
				t.Errorf("Split(%v) = %v; want %v", args, got, want)
			}
		}
	}
}
