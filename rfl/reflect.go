package rfl

import (
	"fmt"
	"reflect"
)

func Fields(arg any) []reflect.StructField {
	t := reflect.TypeOf(arg)
	fs := []reflect.StructField{}
	for _, field := range reflect.VisibleFields(t) {
		if field.Anonymous || !field.IsExported() {
			continue
		}
		fs = append(fs, field)
	}
	return fs
}

func IsField(args ...string) func(reflect.StructField) bool {
	return func(field reflect.StructField) bool {
		for _, arg := range args {
			if field.Name == arg {
				return true
			}
		}
		return false
	}
}

func IsTag(args ...string) func(reflect.StructField) bool {
	return func(field reflect.StructField) bool {
		for _, arg := range args {
			if _, ok := field.Tag.Lookup(arg); ok {
				return true
			}
		}
		return false
	}
}

func SetField(arg any) (func(string, any), error) {
	v := reflect.ValueOf(arg)
	v = v.Elem()
	if !v.CanSet() {
		return nil, fmt.Errorf("struct is not settable")
	}

	return func(key string, val any) {
		field := v.FieldByName(key)
		field.Set(reflect.ValueOf(val))
	}, nil
}
