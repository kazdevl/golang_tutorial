package main

import "fmt"

type Pill int

const (
	A Pill = iota
	B
	C
)

func main() {
	fmt.Printf("%v", B)
}
