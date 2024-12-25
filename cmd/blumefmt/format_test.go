package blumefmt

import "testing"

func TestFormat(t *testing.T) {
	for i, test := range testCases {
		got, err := Fmt([]byte(test.have))
		if err != nil {
			t.Errorf("Test %d failed. err: %s", i, err)
		}
		if string(got) != test.desired {
			t.Errorf("Test %d failed. Got:\n%s\nWant:\n%s\n", i, got, test.desired)
		}
	}
}
