package fileRead

import (
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

func ReadJsonFileContentWithContents() []Content {
	return nil
}

func ReadJsonFileContentWithSliceContent() []string {
	return nil
}
