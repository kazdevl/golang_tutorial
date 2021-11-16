package main

/*
#cgo LDFLAGS: -L. -lhello
#include <hello.sh>
*/

import "C"

// ref: https://qiita.com/hichihara/items/176d4c15bd240d7f2b0d
func main() {
	// cannot exec
	C.hello()
}
