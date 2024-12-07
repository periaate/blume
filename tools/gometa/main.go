package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/periaate/blume/fsio"
	. "github.com/periaate/blume/gen"
	"github.com/periaate/blume/tools/gometa/traits"
	. "github.com/periaate/blume/typ"
)

type Trait struct {
	File    String
	Package string
	Traits  string
	Base    string
	Name    string
}

func (t Trait) Impl() {
	res := traits.Implement(t.Package, t.Name, t.Base, t.Traits)
	fmt.Println(res)
	filename := t.File.ReplaceSuffix(".go", fmt.Sprintf("_%s_impl.go", strings.ToLower(t.Name))).String()
	err := fsio.WriteAll(filename, fsio.B([]byte(res)))
	if err != nil {
		panic(err)
	}
}

func main() {
	traits := []Trait{}
	f := fsio.ReadDirRecursively("./")
	res := Filter(HasSuffix(".go"))(f)
	for _, file := range res {
		str := string(Must(os.ReadFile(file)))
		lines := String(str).Split("\n")
		var Package string
		var derive string
		for _, line := range lines.Array() {
			if line.HasPrefix("package") {
				Package = line.Split(" ").Array()[1].String()
				continue
			}
			if line.HasPrefix("//blume:derive") {
				derive = line.ReplacePrefix("//blume:derive ", "").String()
				continue
			}
			if derive != "" {
				t := line.Split(" ").Array()
				if t[0] != "type" {
					panic("Error: expected type declaration " + line)
				}
				if len(t) < 3 {
					panic("Error: expected type declaration " + line)
				}
				nt := Trait{
					File:    String(file),
					Package: strings.TrimSpace(Package),
					Traits:  strings.TrimSpace(strings.Split(derive, " ")[0]),
					Name:    strings.TrimSpace(t[1].String()),
					Base:    strings.TrimSpace(t[2].String()),
				}
				traits = append(traits, nt)
				derive = ""
			}
		}
	}

	for _, t := range traits {
		t.Impl()
	}
}
