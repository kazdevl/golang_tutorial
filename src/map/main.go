package main

import "fmt"

func main() {
	data := make(map[int]string)
	fmt.Println("len=", len(data)) // mapはcapを使えない
	data = make(map[int]string, 100)
	fmt.Println("len=", len(data)) // mapはcapを使えない
}
