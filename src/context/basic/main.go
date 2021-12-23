package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	c1Ctx := context.WithValue(ctx, "a", "Hoge")
	c2Ctx := context.WithValue(ctx, "b", "Fuga")
	cc1Ctx := context.WithValue(ctx, "c", "HogeHoge")
	// fmt.Println(ctx.Value("a").(string))
	// fmt.Println(ctx.Value("b").(string))
	// fmt.Println(ctx.Value("c").(string))
	fmt.Println(c1Ctx.Value("a").(string))
	// fmt.Println(c2Ctx.Value("a").(string))
	// fmt.Println(cc1Ctx.Value("a").(string))
	// fmt.Println(c1Ctx.Value("b").(string))
	fmt.Println(c2Ctx.Value("b").(string))
	// fmt.Println(cc1Ctx.Value("b").(string))
	// fmt.Println(c1Ctx.Value("c").(string))
	// fmt.Println(c2Ctx.Value("c").(string))
	fmt.Println(cc1Ctx.Value("c").(string))
}
