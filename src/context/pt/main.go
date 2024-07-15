package main

import (
	"fmt"
)

type Sample struct {
	Name string
}

func main() {
	checkValue()
}

func checkValue() {
	v := Sample{Name: "sample"}
	fmt.Printf("before: %p\n", &v)
	defer func() {
		fmt.Printf("defer_アドレス: %p\n", &v)
		fmt.Printf("defer_値: %+v\n", v)
	}()

	v.Name = "changed"
	fmt.Printf("after: %p\n", &v)
}
