package main

import (
	"fmt"
	"io"
	"os"
)

// ref: https://medium.com/eureka-engineering/file-uploads-in-go-with-io-pipe-75519dfa647b
func main() {
	r, w := io.Pipe()
	go func() {
		w.Write([]byte("Hello"))
		fmt.Println("world")
		w.Close()
	}()
	io.Copy(os.Stdout, r)
}
