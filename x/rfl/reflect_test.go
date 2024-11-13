package rfl

import (
	"fmt"
	"testing"

	"github.com/periaate/blume/gen"
)

type Embedded struct {
	Embed [32]byte `json:"abc"`
}

type Test struct {
	privateField string
	PublicField  string `db:"a"`

	TaggedField int `tag:"one" json:"a" db:"b"`
	Embedded
}

func New() Test {
	return Test{
		privateField: "hello world",
		PublicField:  "Hello, World!",
		TaggedField:  20,
	}
}

func TestFields(t *testing.T) {
	val := New()
	f := Fields(val)
	l := len(f)
	if l != 3 {
		t.Error("expected length of 3, got", l)
	}

	for _, field := range f {
		fmt.Println(field.Name)
	}
}

func TestTags(t *testing.T) {
	val := New()
	f := Fields(val)
	f = gen.Filter(IsTag("json"))(f)
	l := len(f)
	if l != 2 {
		t.Error("expected length of 2, got", l)
	}

	for _, field := range f {
		fmt.Println(field.Name)
	}
}

func TestSet(t *testing.T) {
	val := New()

	f, err := SetField(&val)
	if err != nil {
		t.Error(err)
	}

	f("PublicField", "Goodbye, World!")
	if val.PublicField != "Goodbye, World!" {
		t.Error("set field doesn't match, expected \"Goodbye, World!\", received", val.PublicField)
	}
}
