package main

import "fmt"

func main() {
	// use Non-Generic functions
	ints := map[string]int64{
		"1": 10,
		"2": 20,
	}

	floats := map[string]float64{
		"1": 10.01,
		"2": 20.02,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))

	// use Generic function
	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}

// Non-Generic functions
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Generic function
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
