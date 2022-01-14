package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	type Data struct {
		Name  string
		Value string
	}

	d := make([]Data, 100)
	for i := 0; i < len(d); i++ {
		d[i].Name = fmt.Sprintf("SAMPLE%d", i)
		d[i].Value = fmt.Sprintf("%d", i)
	}

	for i := 0; i < 30; i++ {
		basePath, err := filepath.Abs("../../linter/each/")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(basePath)
		outputPath := fmt.Sprintf("%s/%d/main.go", basePath, i)
		if err := generateCode("sample.tpl", outputPath, d); err != nil {
			log.Fatal(err)
		}
	}
}

func generateCode(tmplPath, outputPath string, params interface{}) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	return tmpl.Execute(output, params)
}
