package blob

import (
	"testing"

	"github.com/periaate/blume/gen"
)

func Assert[A any](t *testing.T, msg string, pred gen.Predicate[A]) func(A) {
	return func(a A) {
		if !pred(a) {
			t.Errorf("%s: %v", msg, a)
		}
	}
}

func TestEncode(t *testing.T) {
	Assert(t, "unexpected", gen.Is("AA"))(ContentType(0).Fmt())
	Assert(t, "unexpected", gen.Is("AB"))(ContentType(1).Fmt())
	Assert(t, "unexpected", gen.Isnt("AA"))(ContentType(1).Fmt())

	Assert(t, "unexpected", gen.Is("AA"))(STREAM.Fmt())
	Assert(t, "unexpected", gen.Is("AB"))(PLAIN.Fmt())
}
