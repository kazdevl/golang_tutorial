package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "file/path/sample/sample.txt"
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("not exsit")
	} else {
		fmt.Println("exist")
	}
}
