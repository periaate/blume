package val

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func Sha256[A any](value A) []byte {
	h := sha256.New()
	fmt.Fprint(h, value)
	hash := sha256.Sum256(h.Sum(nil))
	return hash[:]
}

func Sha256R(value io.Reader) []byte {
	h := sha256.New()
	io.Copy(h, value)
	hash := sha256.Sum256(h.Sum(nil))
	return hash[:]
}
