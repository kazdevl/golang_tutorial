package main

import (
	"embed"
	"fmt"
)

//go:embed sample.txt
var sampleStr string

//go:embed sample/sample.html
var sampleHtml string

//go:embed sample/sample.txt
var sampleBytes []byte

//go:embed sample/*
var sampleEmbed embed.FS

//go:embed sample/*.txt sample/nest/*
var sampleFilteredEmbed embed.FS

func main() {
	fmt.Printf("sample.txt=%s\n", sampleStr)
	fmt.Printf("sample/sample.html=%s\n", sampleHtml)
	fmt.Printf("sample/sample.txt(bytes)=%+v\n", sampleBytes)
	fmt.Printf("sample/sample.txt(string)=%s\n", sampleBytes)

	fmt.Println("*******************")
	fmt.Printf("%+v\n", sampleEmbed)
	fs, err := sampleEmbed.ReadDir("sample")
	if err != nil {
		panic(err)
	}
	for _, dir := range fs {
		fmt.Println(dir.Name())
	}

	fmt.Println("*******************")
	fmt.Printf("%+v\n", sampleEmbed)
	fs, err = sampleFilteredEmbed.ReadDir("sample")
	if err != nil {
		panic(err)
	}
	for _, dir := range fs {
		fmt.Println(dir.Name())
	}
	data, err := sampleFilteredEmbed.ReadFile("sample/nest/1/1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("1.txt's data=%s\n", data)
}
