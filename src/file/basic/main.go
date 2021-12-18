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
	f.WriteString("sample1, sample1\n")
	f.WriteString("sample2, sample2\n")

	des, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(des))
	for _, de := range des {
		fmt.Printf("de name(): %s\n", de.Name())
		fsInfo, _ := de.Info()
		fmt.Printf("fsInfo name(): %s\n", fsInfo.Name())
	}
	newPath := filepath.Join(path, "sample")
	fmt.Println(newPath)
	userDir, _ := os.UserHomeDir()
	if err := os.Mkdir(filepath.Join(userDir, "sample"), 0777); err != nil {
		log.Println(err)
	}
	if f, err := os.Stat(filepath.Join(userDir, "sample")); os.IsNotExist(err) || !f.IsDir() {
		fmt.Println("ディレクトリは存在しません！")
	} else {
		fmt.Println("存在するです")
	}
	if err := os.MkdirAll(filepath.Join(userDir, "sample1", "sample2", "sample3"), 0777); err != nil {
		log.Println(err)
	}
	if f, err := os.Stat(filepath.Join(userDir, "sample1", "sample2", "sample3")); os.IsNotExist(err) || !f.IsDir() {
		fmt.Println("ディレクトリは存在しません！")
	} else {
		fmt.Println("存在するです")
	}
}
