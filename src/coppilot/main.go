package main

import (
	"os"
)

func main() {
	if err := CreateFile("sample.txt", "sample"); err != nil {
		panic(err)
	}
	_, err := os.Stat("sample.txt")
	if os.IsNotExist(err) {
		println("not exsit")
	} else {
		println("exist")
	}
}

func CreateFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)
	return nil
}
