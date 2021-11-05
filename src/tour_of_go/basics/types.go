package basics

import "fmt"

type num struct {
	i int
}

// ポインタ
// 値のメモリアドレス
func Pointers() {
	i, j := 10, 20
	p := &i
	fmt.Printf("iの型: %T, pの型(iのポインタ): %T\n", i, p)
	fmt.Printf("iの値: %d, jの値: %d, pの値: %d\n", i, j, *p)
	fmt.Printf("pに代入されたアドレス: %p, iのアドレス: %p\n", p, &i)
	*p = 100
	fmt.Printf("iの値: %d, jの値: %d, pの値: %d\n", i, j, *p)

	num := &num{i: 100}
	fmt.Printf("構造体へのアドレス: %p, numのアドレス: %p, numの値: %+v\n", num, &num, *num)
	pointerArgument(num)
	pointerOfPointerArgument(&num)
}

func pointerArgument(pa *num) {
	// アドレスは値なので、引数の構造体へのアドレスを格納する別のメモリが確保されている
	fmt.Printf("構造体へのアドレス: %p, paのアドレス: %p, paの値: %+v\n", pa, &pa, *pa)
}

func pointerOfPointerArgument(ppa **num) {
	fmt.Printf("num(引数)のアドレス: %p, ppaのアドレス: %p, 構造体へのアドレス: %p, ppaの値: %+v\n", ppa, &ppa, *ppa, **ppa)
}

// 関数値
// 関数を変数に代入することができる
// 関数を引数にすることも可能
func FunctionValue() {
	addFunc := func(x, y int) int {
		return x + y
	}
	fmt.Println(addFunc(1, 1))

	doubleFunc := func(fn func(int, int) int, x, y int) int {
		return fn(x, y) + fn(x, y)
	}
	fmt.Printf("%d\n", doubleFunc(addFunc, 2, 2))
}

// クロージャー
// クロージャーは、それ自身の外部から変数を参照する関数値である。
func FunctionClosures() {
	for i := 0; i < 10; i++ {
		x, y := adder(), adder()
		fmt.Printf("%d番目: x=%d, y=%d\n", i, x(i), y(-i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
