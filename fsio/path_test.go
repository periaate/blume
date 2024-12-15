package fsio_test

import (
	"testing"

	"github.com/periaate/blume/fsio"
	. "github.com/periaate/blume/fsio"
	"github.com/periaate/blume/yap"
	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	res, err := fsio.Join(".", "hi")
	assert.Equal(t, res, "./hi")
	assert.NoError(t, err)
	res, err = fsio.Join("..", "hi")
	assert.Equal(t, res, "../hi")
	assert.NoError(t, err)
	res, err = fsio.Join("base", "hi")
	assert.Equal(t, res, "base/hi")
	assert.NoError(t, err)
	res, err = fsio.Join("base/", "hi")
	assert.Equal(t, res, "base/hi")
	assert.NoError(t, err)

	res, err = fsio.Join("./", "hi")
	assert.Equal(t, res, "./hi")
	assert.NoError(t, err)
	res, err = fsio.Join("../", "hi")
	assert.Equal(t, res, "../hi")
	assert.NoError(t, err)
	res, err = fsio.Join("./base", "hi")
	assert.Equal(t, res, "./base/hi")
	assert.NoError(t, err)
	res, err = fsio.Join("./base/", "hi")
	assert.Equal(t, res, "./base/hi")
	assert.NoError(t, err)

	res, err = fsio.Join("./", "/hi")
	assert.Equal(t, res, "./hi")
	assert.NoError(t, err)
	res, err = fsio.Join("../", "/hi")
	assert.Equal(t, res, "../hi")
	assert.NoError(t, err)
	res, err = fsio.Join("./base", "/hi")
	assert.Equal(t, res, "./base/hi")
	assert.NoError(t, err)
	res, err = fsio.Join("./base/", "/hi")
	assert.Equal(t, res, "./base/hi")
	assert.NoError(t, err)

	res, err = fsio.Join("./", "hi", "world")
	assert.Equal(t, res, "./hi/world")
	assert.NoError(t, err)
	res, err = fsio.Join("../", "hi", "world")
	assert.Equal(t, res, "../hi/world")
	assert.NoError(t, err)
	res, err = fsio.Join("./base", "hi", "world")
	assert.Equal(t, res, "./base/hi/world")
	assert.NoError(t, err)
	res, err = fsio.Join("./base/", "hi", "world")
	assert.Equal(t, res, "./base/hi/world")
	assert.NoError(t, err)

	res, err = fsio.Join("./", "//hi")
	assert.Equal(t, res, "./hi")
	assert.NoError(t, err)
	res, err = fsio.Join("../", "//hi")
	assert.Equal(t, res, "../hi")
	assert.NoError(t, err)
	res, err = fsio.Join("./base", "//hi")
	assert.Equal(t, res, "./base/hi")
	assert.NoError(t, err)
	res, err = fsio.Join("./base/", "//hi")
	assert.Equal(t, res, "./base/hi")
	assert.NoError(t, err)

	res, err = fsio.Join("./", "hi/", "//world")
	assert.Equal(t, res, "./hi/world")
	assert.NoError(t, err)
	res, err = fsio.Join("../", "hi//", "//world")
	assert.Equal(t, res, "../hi/world")
	assert.NoError(t, err)

	res, err = fsio.Join("./", "hi/", "//world/")
	assert.Equal(t, res, "./hi/world/")
	assert.NoError(t, err)
	res, err = fsio.Join("../", "hi//", "//world/")
	assert.Equal(t, res, "../hi/world/")
	assert.NoError(t, err)

	res, err = fsio.Join("~/", "hi/", "//world/")
	assert.Equal(t, res, "~/hi/world/")
	assert.NoError(t, err)
	res, err = fsio.Join("/", "hi//", "//world/")
	assert.Equal(t, res, "/hi/world/")
	assert.NoError(t, err)

	res, err = fsio.Join(`~\`, `hi\`, `\\world\`)
	assert.Equal(t, res, `~/hi/world/`)
	assert.NoError(t, err)
	res, err = fsio.Join(`\`, `hi\\`, `\\world\`)
	assert.Equal(t, res, `/hi/world/`)
	assert.NoError(t, err)

	res, err = fsio.Join(`/`)
	assert.Equal(t, res, `/`)
	assert.NoError(t, err)
	res, err = fsio.Join(`\`)
	assert.Equal(t, res, `/`)
	assert.NoError(t, err)
	res, err = fsio.Join(`./`)
	assert.Equal(t, res, `./`)
	assert.NoError(t, err)
	res, err = fsio.Join(`.\`)
	assert.Equal(t, res, `./`)
	assert.NoError(t, err)

	res, err = fsio.Join(``)
	assert.Error(t, err)
	res, err = fsio.Join(`.`)
	assert.Equal(t, res, `./`)
	assert.NoError(t, err)

	res, err = fsio.Join(`./blob/`, `test/AAAA`)
	assert.Equal(t, res, `./blob/test/AAAA`)
	assert.NoError(t, err)
	res, err = fsio.Join(`http://127.0.0.1:8085`, `b`, `./`, `test/AAAAAAAAAAAAAAAAAAA`)
	assert.Equal(t, res, `http://127.0.0.1:8085/b/./test/AAAAAAAAAAAAAAAAAAA`)
	assert.NoError(t, err)

	res, err = fsio.Join(`http://`, `0.0.0.0:8000`, `/`)
	assert.Equal(t, res, `http://0.0.0.0:8000/`)
	assert.NoError(t, err)
	res, err = fsio.Join(`http://`, `//0.0.0.0:8000/`, `//`)
	assert.Equal(t, res, `http://0.0.0.0:8000/`)
	assert.NoError(t, err)
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

	yap.SetLevel(yap.L_Debug)
	for _, tc := range testCases {
		t.Run(tc.inp, func(t *testing.T) {
			res := Clean(tc.inp)
			assert.Equal(t, tc.exp, res)
		})
	}
}
