package fsio_test

import (
	"path/filepath"
	"testing"

	. "github.com/periaate/blume/fsio"
	"github.com/periaate/blume/pred/has"
	"github.com/periaate/blume/pred/is"
	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	assert.True(t, is.Ok(First("./test", func(path string) bool {
		return has.Prefix("hii")(filepath.Base(path))
	})))
}

func TestFind(t *testing.T) {
	assert.Len(t, Find("./test", func(path string) bool {
		return has.Any(".test")(filepath.Base(path))
	}), 3)
}
