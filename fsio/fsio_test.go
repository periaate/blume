package fsio

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/periaate/blume/clog"
)

type PathCase struct {
	Elems    []string
	Expected string
}

var TestCases = []PathCase{
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
}

func TestJoin(t *testing.T) {
	clog.SetLogLoggerLevel(clog.LevelDebug)
	for _, tc := range TestCases {
		t.Run(strings.Join(tc.Elems, "/"), func(t *testing.T) {
			res := Join(tc.Elems...)
			if res != tc.Expected {
				clog.Error("unexpcted result", "res", res, "expected", tc.Expected)
				t.Fail()
			}
			clog.Debug("comparison", "res", res, "filepath", filepath.Join(tc.Elems...))
		})
	}
}
