package blush

import (
	"fmt"
	"os"
	"testing"

	"github.com/periaate/blume/clog"
	"github.com/periaate/blume/fsio"
)

func openCases() (res []string, err error) {
	res, err = fsio.ReadDir("./tests")
	if err != nil {
		return
	}

	var ress []string

	for _, v := range res {
		fmt.Println(v)
		f, err := os.ReadFile(v)
		if err != nil {
			return res, err
		}

		fmt.Println(string(f))
		ress = append(ress, string(f))
	}

	return ress, nil
}

func TestEval(t *testing.T) {
	clog.SetLogLoggerLevel(clog.LevelDebug)
	testCases, err := openCases()
	if err != nil {
		return
	}

	if len(testCases) == 0 {
		t.Fatal("length of test cases is 0")
	}

	for _, v := range testCases {
		res, err := Eval(v)
		if err != nil {
			return
		}

		fmt.Println(res)
	}
}
