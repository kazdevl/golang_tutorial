package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("./../testdata")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
	f, _ := os.Create(fmt.Sprintf("%s/sample1.txt", path))
	defer f.Close()
	f.Write([]byte("sample22"))
}
