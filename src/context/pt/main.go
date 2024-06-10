package main

import (
	"context"
	"fmt"
)

type Sample struct {
	Name string
}

func main() {
	// ctx := context.Background()
	// defer func() {
	// 	s := ctx.Value("sample").(*Sample)
	// 	println(s.Name)
	// }()

	// s := &Sample{Name: "sample"}
	// ctx = context.WithValue(context.Background(), "sample", s)
	// nothing(ctx)

	// checkPointer()
	checkValue()
}

func nothing(_ context.Context) {
}

func checkPointer() {
	v := &Sample{Name: "sample"}
	fmt.Printf("before: %p\n", v)
	defer func() {
		fmt.Printf("defer_アドレス: %p\n", v)
		fmt.Printf("defer_値: %+v\n", v)
	}()

	v.Name = "changed"
	fmt.Printf("after: %p\n", v)
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
