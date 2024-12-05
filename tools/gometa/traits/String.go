package traits

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"text/template"
)

//go:embed String.trait.tmpl
var StringTrait string

func s(name, base string) string {
	fmt.Println("String Impl")
	fmt.Println(name)
	fmt.Println(base)
	fmt.Println(base == "string")
	if base == "" || name == "" {
		log.Fatalf("Error: name and base must be provided")
	}

	if base != "string" {
		log.Fatalf("Error: base type must be string")
	}

	tmpl, err := template.New(name).Parse(StringTrait)
	if err != nil {
		log.Fatalf("Error parsing trait template (%s): %v", name, err)
	}

	var b bytes.Buffer

	if err := tmpl.Execute(&b, name); err != nil {
		log.Fatalf("Error executing trait template (%s): %v", name, err)
	}

	str := b.String()
	if str == "" {
		log.Fatalf("Error: empty trait template (%s)", name)
	}

	return str
}
