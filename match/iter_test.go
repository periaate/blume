package match

import (
	"fmt"
	"testing"
)

func TestIter(t *testing.T) {
	func() {
		var i int
		var v byte
		itr, err := ToIter[string, byte]("0123456789")
		if err != nil { t.Error(err); return }
		for i, v = range itr.Iter() {
			if fmt.Sprint(i) != string(v) { t.Errorf("%v != %s", i, string(v)) }
		}
		if i != 9 { t.Error("string|byte iterator didn't iterate 10 times") }
	}()

	func() {
		itr, err := ToIter[string, string]("0123456789")
		if err != nil { t.Error(err); return }
		var i int
		var v string
		for i, v = range itr.Iter() {
			if fmt.Sprint(i) != v { t.Errorf("%v != %s", i, v) }
		}
		if i != 9 { t.Error("string|string iterator didn't iterate 10 times") }
	}()
}
