package fileRead

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Content struct {
	Page string `json:"page"`
}

type SliceContent struct {
	Pages []string `json:"pages"`
}

func ReadTextFileContent() []string {
	data, err := os.ReadFile("sample1.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

func ReadJsonFileContentWithContents() []string {
	f, _ := os.Open("sample1.json")
	var data []Content
	json.NewDecoder(f).Decode(&data)
	strs := make([]string, len(data))
	for i, v := range data {
		strs[i] = v.Page
	}
	return strs
}

func ReadJsonFileContentWithSliceContent() []string {
	f, _ := os.Open("sample2.json")
	var data SliceContent
	json.NewDecoder(f).Decode(&data)
	return data.Pages
}
