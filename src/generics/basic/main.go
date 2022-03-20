package main

import "fmt"

// type constraint
type Number interface {
	int64 | float64
}
type UnderlyingNumber interface {
	~int64 | ~float64
}

type Sample int64
type SSample Sample

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
	fmt.Printf("Generic Sums with Constraint: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))

	// this code can't pass constraint
	underlyingIntsVer1 := map[string]Sample{
		"1": Sample(1),
		"2": Sample(2),
	}
	underlyingIntsVer2 := map[string]SSample{
		"1": SSample(1),
		"2": SSample(2),
	}
	// fmt.Printf("Sample and SSample don'nt implement Number %v and %v\n", SumIntsOrFloats(underlyingIntsVer1), SumIntsOrFloats(underlyingIntsVer2))
	fmt.Printf("Sample and SSample implement UnderlyingNumber %v and %v\n", SumUnderlyingIntsOrFloats(underlyingIntsVer1), SumUnderlyingIntsOrFloats(underlyingIntsVer2))

	// use 1 type for 2 arguments
	fmt.Print(SubIntsOrFloats[int64, float64](int64(10), 11.2))
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
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumUnderlyingIntsOrFloats[K comparable, V UnderlyingNumber](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SubIntsOrFloats[A Number, B Number](a A, b B) A {
	return A(int(a) - int(b))
}
