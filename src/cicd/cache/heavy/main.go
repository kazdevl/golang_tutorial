package main

import (
	"fmt"
	"math"

	"github.com/kazdevl/golang_tutorial/cicd/cache/heavy/greet"
)

func main() {
	greet.CallGreetAll()

	fmt.Println(heavyProcess(100))
}

func heavyProcess(n int) float64 {
	if n == 1 {
		return 1
	}
	// This function is used to simulate a heavy process.
	return math.Sqrt(float64(n) * heavyProcess(n-1))
}
