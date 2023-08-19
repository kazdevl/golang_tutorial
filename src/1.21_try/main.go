package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	// min・max
	fmt.Println("min・maxの動作確認")
	fmt.Println(min(1, 2, 3, 4, 5))
	fmt.Println(max(11, 13, 14, 9))
	fmt.Println(max("", "a", "c", "d"))
	fmt.Println(min("", "a", "c", "d"))

	// cmpパッケージの動作確認
	fmt.Println("cmp.Compareの動作確認")
	fmt.Println(cmp.Compare[float64](0.0, 1.2))
	fmt.Println(cmp.Compare[float64](1.2, 0.0))
	fmt.Println(cmp.Compare[float64](1.2, 1.2))

	fmt.Println("cmp.Lessの動作確認")
	fmt.Println(cmp.Less[float64](0.0, 1.2))
	fmt.Println(cmp.Less[float64](1.2, 0.0))

	// slicesパッケージの動作確認
	fmt.Println("slicesの動作確認")
	ss := make([]int, 0, 10)
	ss = append(ss, 1, 2, 3, 4, 5)
	slices.Reverse(ss)
	fmt.Println("reversed", ss)
	ss = slices.Replace(ss, 1, 3, 10, 20)
	fmt.Println("replaced", ss)
	ss = slices.Delete(ss, 1, 3)
	fmt.Println("deleted", ss)
	ss = slices.Insert(ss, 1, 100, 200)
	fmt.Println("inserted", ss)
}
