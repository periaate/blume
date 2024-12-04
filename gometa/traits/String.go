package traits

import (
	"bytes"
	_ "embed"
	"log"
	"text/template"
)

//go:embed String.trait.tmpl
var StringTrait string

func String(name, base string) string {
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

	return b.String()
}
