package main

import (
	"log"
	"os"
	"text/template"
)

type Sample struct {
	A string
	B int
}

func main() {
	sample := Sample{"Hello", 2}
	tmpl, err := template.New("sample").Parse("Aの背丈は{{.A}}. B is {{.B}}")
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.Execute(os.Stdout, sample); err != nil {
		log.Fatal(err)
	}
}
