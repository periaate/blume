package media

import (
	"encoding/hex"
	"os"
	"path"
	"strings"
	"testing"
)

const (
	// Expected hash as hex string of all non-fail images in ./test
	expect = "3e7c3dbc3bdc37ec2ff41ff8c00340024002c0031ff82ff437ec3bdc3dbc3e7c"
)

// test hashes of all files in ./test
// MakePhash also calls OpenImage, which tests thumbnailing indirectly
func TestPerceptualHash(t *testing.T) {
	files, err := os.ReadDir("./test")
	if err != nil {
		t.Error(err)
	}

	for _, file := range files {
		file, err := os.Open(path.Join("./test", file.Name()))
		if err != nil {
			t.Fatal("error opening test file", err)
		}

		img, err := ImageFromReader(file)
		if err != nil {
			t.Fatal("error reading file to image", err)
		}

		hash := GeneratePhash(img, 16)

		shouldFail := strings.Contains(file.Name(), "fail")
		if hex.EncodeToString(hash[:]) != expect && !shouldFail {
			t.Errorf("expected %s, got %s", expect, hex.EncodeToString(hash))
		}
	}
}
