package main

import "fmt"

func main() {
	// min・max
	fmt.Println(min(1, 2, 3, 4, 5))
	fmt.Println(max(11, 13, 14, 9))
	fmt.Println(max("", "a", "c", "d"))
	fmt.Print(min("", "a", "c", "d"))
}
