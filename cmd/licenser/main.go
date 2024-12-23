package main

import (
	_ "embed"
	"fmt"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/yap"
)

func main() {
	yap.IncludeTimes(false, false, false, false, false, false)

	if Any(Is("LICENSE", "License", "license"))(Must(fsio.ReadDir("./")).Val) {
		yap.Fatal("license file already exists")
	}

	s := fsio.Args[string](func(s []string) bool { return len(s) == 1 }).Must().Val
	fmt.Printf("running licenser with argument %s\n", s[0])
	if lic, ok := licenses[s[0]]; ok {
		err := fsio.WriteNew("LICENSE", Buf(lic))
		if err != nil {
			yap.Fatal("error writing license", "err", err)
		}
		fmt.Printf("License %s written to LICENSE\n", s[0])
	} else {
		fmt.Printf("License %s not found\n\nAvailable licenses:\n%s", s[0], licensesStr)
	}
}

var licensesStr = `AGPL-3.0
GPL-3.0
GPL-2.0
LGPL-3.0
LGPL-2.1
NON-AI-MPL-2.0
MPL-2.0
NON-AI-APACHE-2.0
Apache-2.0
NON-AI-UNLICENSE
UNLICENSE
NON-AI-MIT
MIT`

var licenses = map[string][]byte{
	"GPL-2.0":           gpl20,
	"GPL-3.0":           gpl30,
	"AGPL-3.0":          agpl30,
	"MIT":               mit,
	"MPL-2.0":           mpl20,
	"Apache-2.0":        apache20,
	"LGPL-2.1":          lgpl21,
	"LGPL-3.0":          lgpl30,
	"UNLICENSE":         unlicense,
	"NON-AI-MIT":        non_ai_mit,
	"NON-AI-MPL-2.0":    non_ai_mpl20,
	"NON-AI-UNLICENSE":  non_ai_unlicense,
	"NON-AI-APACHE-2.0": non_ai_apache20,
}

//go:embed GPL-2.0
var gpl20 []byte

//go:embed GPL-3.0
var gpl30 []byte

//go:embed AGPL-3.0
var agpl30 []byte

//go:embed MIT
var mit []byte

//go:embed MPL-2.0
var mpl20 []byte

//go:embed Apache-2.0
var apache20 []byte

//go:embed LGPL-2.1
var lgpl21 []byte

//go:embed LGPL-3.0
var lgpl30 []byte

//go:embed UNILICENSE
var unlicense []byte

//go:embed NON-AI-MIT
var non_ai_mit []byte

//go:embed NON-AI-MPL-2.0
var non_ai_mpl20 []byte

//go:embed NON-AI-UNLICENSE
var non_ai_unlicense []byte

//go:embed NON-AI-APACHE-2.0
var non_ai_apache20 []byte
