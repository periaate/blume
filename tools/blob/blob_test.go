package blob

import (
	"bytes"
	"testing"
)

func TestPath(t *testing.T) {
	SetIndex("./testing")
	exp := []struct {
		Inp    string
		Type   ContentType
		File   string
		Bucket string
		Blob   string
		Err    bool
	}{
		{"a/b", PLAIN, "./testing/a/" + PLAIN.Fmt() + "b", "a", "b", false},
		{"a/b/", PLAIN, "./testing/a/" + PLAIN.Fmt() + "b", "a", "b", false},
		{"/a/b", PLAIN, "./testing/a/" + PLAIN.Fmt() + "b", "a", "b", false},
		{"/a/b/", PLAIN, "./testing/a/" + PLAIN.Fmt() + "b", "a", "b", false},
	}

	for _, e := range exp {
		b := Blob(e.Inp)
		if !e.Err {
			I.Set(b, e.Type)
		}

		bucket, blob, err := b.Split()
		if e.Err {
			if err == nil {
				t.Errorf("Expected error, got nil: %s", err)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if bucket != e.Bucket {
			t.Errorf("Expected bucket %s, got %s", e.Bucket, bucket)
		}

		if blob != e.Blob {
			t.Errorf("Expected blob %s, got %s", e.Blob, blob)
		}

		file, ct, err := b.File()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if ct != e.Type {
			t.Errorf("Expected type %v, got %v", e.Type, ct)
		}

		if file != e.File {
			t.Errorf("Expected file %s, got %s", e.File, file)
		}

		ct, err = b.Type()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if ct != e.Type {
			t.Errorf("Expected type %v, got %v", e.Type, ct)
		}
	}

	CloseIndex()
}

func TestIndex(t *testing.T) {
	SetIndex("./testing")
	exp := []struct {
		Inp   string
		Type  ContentType
		Value string
	}{
		{"a/b", PLAIN, "Hello World!"},
		{"a/b", HTML, "<html><body>Hello World!</body></html>"},
		{"a/b", JSON, `{"msg":"Hello World!"}`},

		{"/a/b", PLAIN, "Hello World!"},
		{"/a/b", HTML, "<html><body>Hello World!</body></html>"},
		{"/a/b", JSON, `{"msg":"Hello World!"}`},

		{"a/b/", PLAIN, "Hello World!"},
		{"a/b/", HTML, "<html><body>Hello World!</body></html>"},
		{"a/b/", JSON, `{"msg":"Hello World!"}`},

		{"/a/b//", PLAIN, "Hello World!"},
		{"/a//b/", HTML, "<html><body>Hello World!</body></html>"},
		{"//a/b/", JSON, `{"msg":"Hello World!"}`},
	}

	for _, e := range exp {
		b := Blob(e.Inp)

		err := b.Set(bytes.NewBufferString(e.Value), e.Type)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}

		r, ct, err := b.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}

		if ct != e.Type {
			t.Errorf("Expected type %v, got %v", e.Type, ct)
			continue
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		if buf.String() != e.Value {
			t.Errorf("Expected value %s, got %s", e.Value, buf.String())
			continue
		}

		err = b.Del()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}
	}

	CloseIndex()
}

func TestErrors(t *testing.T) {
	SetIndex("./testing")
	exp := []struct {
		Inp   string
		Type  ContentType
		Value string
	}{
		{"./a/b", PLAIN, "Hello World!"},
		{"a/b/.", HTML, "<html><body>Hello World!</body></html>"},
	}

	for _, e := range exp {
		b := Blob(e.Inp)

		err := b.Set(bytes.NewBufferString(e.Value), e.Type)
		if err == nil {
			t.Errorf("Expected error, got nil")
			continue
		}
	}
}
