package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type Data struct {
	PackageName string
	TypeName    string
	BaseType    string
	TraitNames  []string
}

//go:embed traits/*
var traits embed.FS

func main() {
	if len(os.Args) < 6 || strings.ToLower(os.Args[3]) != "as" || strings.ToLower(os.Args[5]) != "derive" {
		log.Fatalf("Usage: //go:generate gometa <PackageName> <TypeName> as <BaseType> derive <TraitName1> <TraitName2> ...")
	}

	d := Data{
		PackageName: os.Args[1],
		TypeName:    os.Args[2],
		BaseType:    os.Args[4],
		TraitNames:  os.Args[6:],
	}

	outputName := fmt.Sprintf("%s_impl.go", strings.ToLower(d.TypeName))

	funcMap := template.FuncMap{
		"impl": impl,
		"assert": func(a, b string) {
			if a != b {
				log.Fatalf("Assertion failed: %s != %s", a, b)
			}
		},
	}

	tmpl, err := template.New("base").Funcs(funcMap).ParseFS(traits, "traits/base.tmpl")
	if err != nil {
		log.Fatalf("Error parsing base template: %v", err)
	}

	f, err := os.Create(outputName)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, d); err != nil {
		log.Fatalf("Error executing base template: %v", err)
	}

	log.Printf("Generated %s", outputName)
}

func impl(typeName string, baseType string, traitNames ...string) string {
	var res bytes.Buffer
	for _, t := range traitNames {
		tmplPath := fmt.Sprintf("traits/%s.trait.tmpl", t)
		tmpl, err := template.New(t).ParseFS(traits, tmplPath)
		if err != nil {
			log.Fatalf("Error parsing trait template (%s): %v", tmplPath, err)
		}

		var b bytes.Buffer
		d := Data{TypeName: typeName, BaseType: baseType}

		if err := tmpl.Execute(&b, d); err != nil {
			log.Fatalf("Error executing trait template (%s): %v", tmplPath, err)
		}

		res.Write(b.Bytes())
	}
	return res.String()
}
