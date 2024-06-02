package main

import (
	"fmt"
	"math"

	"github.com/kazdevl/golang_tutorial/cicd/cache/heavy/greet"
)

func main() {
	greet.CallGreetAll()

	fmt.Println(heavyProcess(1000))
}

func heavyProcess(n int) float64 {
	// This function is used to simulate a heavy process.
	return math.Sqrt(float64(n) * heavyProcess(n-1))
}
