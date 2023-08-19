package main

import (
	"cmp"
	"fmt"
)

func main() {
	// min・max
	fmt.Println("min・maxの動作確認")
	fmt.Println(min(1, 2, 3, 4, 5))
	fmt.Println(max(11, 13, 14, 9))
	fmt.Println(max("", "a", "c", "d"))
	fmt.Println(min("", "a", "c", "d"))

	// cmp
	fmt.Println("cmp.Compareの動作確認")
	fmt.Println(cmp.Compare[float64](0.0, 1.2))
	fmt.Println(cmp.Compare[float64](1.2, 0.0))
	fmt.Println(cmp.Compare[float64](1.2, 1.2))

	fmt.Println("cmp.Lessの動作確認")
	fmt.Println(cmp.Less[float64](0.0, 1.2))
	fmt.Println(cmp.Less[float64](1.2, 0.0))

}
