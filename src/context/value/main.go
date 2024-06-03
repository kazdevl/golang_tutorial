package main

import (
	"context"
	"fmt"

	"github.com/kazdevl/golang_tutorial/context/value/econtext"
	"github.com/kazdevl/golang_tutorial/context/value/rcontext"
)

func main() {
	ctx := context.Background()
	ctx = econtext.Set(ctx)
	ctx = rcontext.Set(ctx)

	if ectx := econtext.Extract(ctx); ectx != nil {
		fmt.Printf("ectx = %+v\n", ectx)
	} else {
		fmt.Println("ectx is nil")
	}
	if rctx := rcontext.Extract(ctx); rctx != nil {
		fmt.Printf("rctx = %+v\n", rctx)
	} else {
		fmt.Println("rctx is nil")
	}
}
