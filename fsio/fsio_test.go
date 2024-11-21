package fsio

import (
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/periaate/blume/clog"
)

func TestJoin(t *testing.T) {
	TestCases := []struct {
		Elems    []string
		Expected string
	}{
		{[]string{".", "hi"}, "./hi"},
		{[]string{"..", "hi"}, "../hi"},
		{[]string{"base", "hi"}, "base/hi"},
		{[]string{"base/", "hi"}, "base/hi"},

		{[]string{"./", "hi"}, "./hi"},
		{[]string{"../", "hi"}, "../hi"},
		{[]string{"./base", "hi"}, "./base/hi"},
		{[]string{"./base/", "hi"}, "./base/hi"},

		{[]string{"./", "/hi"}, "./hi"},
		{[]string{"../", "/hi"}, "../hi"},
		{[]string{"./base", "/hi"}, "./base/hi"},
		{[]string{"./base/", "/hi"}, "./base/hi"},

		{[]string{"./", "hi", "world"}, "./hi/world"},
		{[]string{"../", "hi", "world"}, "../hi/world"},
		{[]string{"./base", "hi", "world"}, "./base/hi/world"},
		{[]string{"./base/", "hi", "world"}, "./base/hi/world"},

		{[]string{"./", "//hi"}, "./hi"},
		{[]string{"../", "//hi"}, "../hi"},
		{[]string{"./base", "//hi"}, "./base/hi"},
		{[]string{"./base/", "//hi"}, "./base/hi"},

		{[]string{"./", "hi/", "//world"}, "./hi/world"},
		{[]string{"../", "hi//", "//world"}, "../hi/world"},

		{[]string{"./", "hi/", "//world/"}, "./hi/world/"},
		{[]string{"../", "hi//", "//world/"}, "../hi/world/"},

		{[]string{"~/", "hi/", "//world/"}, "~/hi/world/"},
		{[]string{"/", "hi//", "//world/"}, "/hi/world/"},

		{[]string{`~\`, `hi\`, `\\world\`}, `~/hi/world/`},
		{[]string{`\`, `hi\\`, `\\world\`}, `/hi/world/`},

		{[]string{`/`}, `/`},
		{[]string{`\`}, `/`},
		{[]string{`./`}, `./`},
		{[]string{`.\`}, `./`},

		{[]string{``}, ``},
		{[]string{`.`}, `./`},

		{[]string{`./blob/`, `test/AAAA`}, `./blob/test/AAAA`},
		{
			[]string{`http://127.0.0.1:8085`, `b`, `./`, `test/AAAAAAAAAAAAAAAAAAA`},
			`http://127.0.0.1:8085/b/./test/AAAAAAAAAAAAAAAAAAA`,
		},

		{[]string{`http://`, `0.0.0.0:8000`, `/`}, `http://0.0.0.0:8000/`},
		{[]string{`http://`, `//0.0.0.0:8000/`, `//`}, `http://0.0.0.0:8000/`},
	}

	clog.SetLogLoggerLevel(clog.LevelDebug)
	for _, tc := range TestCases {
		t.Run(strings.Join(tc.Elems, "/"), func(t *testing.T) {
			res := Join(tc.Elems...)
			if res != tc.Expected {
				clog.Error("unexpcted result", "res", res, "expected", tc.Expected)
				t.Fail()
			}
			clog.Debug("comparison", "res", res, "filepath", filepath.Join(tc.Elems...), "path", path.Join(tc.Elems...))
		})
	}
}

func TestClean(t *testing.T) {
	testCases := []struct {
		inp string
		exp string
	}{
		{"base//clean", "base/clean"},
		{`base\\clean`, "base/clean"},

		{"http://base//clean", "http://base/clean"},
		{"http:///base//clean", "http://base/clean"},
	}

	clog.SetLogLoggerLevel(clog.LevelDebug)
	for _, tc := range testCases {
		t.Run(tc.inp, func(t *testing.T) {
			res := Clean(tc.inp)
			if res != tc.exp {
				clog.Error("unexpcted result", "res", res, "expected", tc.exp)
				t.Fail()
			}
			clog.Debug("comparison", "res", res, "filepath", filepath.Clean(tc.inp))
		})
	}
}
