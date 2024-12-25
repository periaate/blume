package num

import (
	"testing"
)

func TestClamp(t *testing.T) {
	cases := []struct {
		inp int
		min int
		max int
		exp int
	}{
		{5, 0, 10, 5},
		{0, 0, 10, 0},
		{-5, 0, 10, 0},
		{-5, -10, 10, -5},
		{-5, 10, -10, -5},
		{-50, -10, 10, -10},
		{-50, 10, -10, -10},
		{5, -10, 10, 5},
		{5, 10, -10, 5},
		{50, -10, 10, 10},
		{50, 10, -10, 10},
	}

	for _, c := range cases {
		got := Clamp(c.min, c.max)(c.inp)
		if got != c.exp {
			t.Errorf("SmartClamp(%d, %d) == %d, want %d", c.inp, c.max, got, c.exp)
		}
	}
}

func TestAbs(t *testing.T) {
	cases := []struct {
		inp int
		exp int
	}{
		{inp: 5, exp: 5},
		{inp: -5, exp: 5},
		{inp: 0, exp: 0},
	}

	for _, c := range cases {
		got := Abs(c.inp)
		if got != c.exp {
			t.Errorf("Abs(%d) == %d, want %d", c.inp, got, c.exp)
		}
	}
}
