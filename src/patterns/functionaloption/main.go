package main

import (
	"fmt"

	"github.com/kazdevl/golang_tutorial/patterns/functionaloption/option"
)

func main() {
	fmt.Printf("result with option: %v\n", option.NewExampleOption(option.WithID(10), option.WithClient("sample")))
	fmt.Printf("result without option: %v\n", option.NewExampleOption())
}
