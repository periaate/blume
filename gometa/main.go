package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/periaate/blume/fsio"
	. "github.com/periaate/blume/gen"
	"github.com/periaate/blume/gometa/traits"
)

type Trait struct {
	File    String
	Package string
	Traits  []string
	Base    string
	Name    string
}

func (t Trait) Impl() {
	res := traits.Implement(t.Package, t.Name, t.Base, t.Traits...)
	filename := t.File.ReplaceSuffix(".go", fmt.Sprintf("_%s_impl.go", strings.ToLower(t.Name))).String()
	err := fsio.WriteAll(filename, fsio.B([]byte(res)))
	if err != nil {
		panic(err)
	}
}

func main() {
	traits := []Trait{}

	res := Filter(HasSuffix(".go"))(fsio.ReadDirRecursively("./"))
	for _, file := range res {
		str := string(Must(os.ReadFile(file)))
		lines := String(str).Split("\n")
		var Package string
		var derive string
		for _, line := range lines {
			if line.HasPrefix("package") {
				Package = line.Split(" ")[1].String()
				continue
			}
			if line.HasPrefix("//blume:derive") {
				derive = line.ReplacePrefix("//blume:derive ").String()
				continue
			}
			if derive != "" {
				t := line.Split(" ")
				if t[0] != "type" {
					panic("Error: expected type declaration " + line)
				}
				if len(t) < 3 {
					panic("Error: expected type declaration " + line)
				}
				traits = append(traits, Trait{
					File:    String(file),
					Package: Package,
					Traits:  strings.Split(derive, " "),
					Name:    t[1].String(),
					Base:    t[2].String(),
				})
			}
		}
	}

	for _, t := range traits {
		t.Impl()
	}
}

// fmt.Println(strings.Join(os.Args, " "))
// 	if len(os.Args) < 6 || strings.ToLower(os.Args[4]) != "as" || strings.ToLower(os.Args[6]) != "derive" {
// 		log.Fatalf("Usage: //go:generate gometa <RelPath> <PackageName> <TypeName> as <BaseType> derive <TraitName1> <TraitName2> ...")
// 	}
//
// 	RelPath := os.Args[1]
// 	PackageName := os.Args[2]
// 	TypeName := os.Args[3]
// 	BaseType := os.Args[5]
// 	traitNames := os.Args[7:]
//
// 	fmt.Println(PackageName, TypeName, BaseType, traitNames)
//
// 	outputName := filepath.Join(RelPath, fmt.Sprintf("%s_impl.go", strings.ToLower(TypeName)))
//
// 	res := traits.Implement(PackageName, TypeName, BaseType, traitNames...)
// 	if res == "" {
// 		log.Fatalf("Error: no implementation generated")
// 	}
//
// 	err := fsio.WriteAll(outputName, fsio.B([]byte(res)))
// 	if err != nil {
// 		log.Fatalf("Error writing implementation to file: %v", err)
// 	}
//
// 	fmt.Printf("Generated implementation(s) for %s in %s\n", TypeName, outputName)
// }
