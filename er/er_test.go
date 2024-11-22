package er

import (
	"net/http"
	"testing"
)

func Assert[K comparable](t *testing.T, has, expects K) {
	if has != expects {
		t.Errorf("assertion failed! expected: [%v], got: [%v]", expects, has)
	}
}

func TestErrors(t *testing.T) {
	Assert(t, BadRequest{}.Status(), http.StatusBadRequest)
	Assert(t, NotFound{}.Status(), http.StatusNotFound)
}
