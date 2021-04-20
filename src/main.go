package main

import (
	"app/basics"
	"fmt"
)

func main() {
	// 参考url: https://go-tour-jp.appspot.com/basics/11
	checkBasic()
}

func checkBasic() {
	// 1
	fmt.Println(basics.Add(10, 20))
	// 2
	a := "Sing"
	b := "Song"
	a, b = basics.Swap(a, b)
	fmt.Println(a, b)
	// 3
	basics.Variables()
	// 4
	basics.VariablesWithInitializers()
	// 5
	basics.ShortVariableDeclarations()
	// 6
	basics.TypeConversions()
	// 7
	basics.Constants()
}
