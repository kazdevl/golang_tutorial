package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("start")
	for {
		time.Sleep(100 * time.Millisecond)
		fib(30)
	}
}
