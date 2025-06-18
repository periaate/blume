package test

import (
	"reflect"
	"testing"
)

type Test struct {
	*testing.T
	res any
}

func Expect(t *testing.T, v any) *Test { return &Test{ t, v }}

func (t *Test) Is(v any, args ...any) {
	if !reflect.DeepEqual(t.res, v) {
		t.Error(args...)
	}
}

func (t *Test) IsF(v any, f string, args ...any) {
	if !reflect.DeepEqual(t.res, v) {
		t.Errorf(f, args...)
	}
}

func ErrIf(t *testing.T, v bool, f string, args ...any) {
	if v { t.Errorf(f, args...) }
}
