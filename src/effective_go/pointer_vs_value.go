package effective_go

import (
	"fmt"
)

// 参考文献1: https://golang.org/doc/effective_go#methods
// 参考文献2: https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/#u30B3_u30F3_u30D1_u30A4_u30E9_u306E_u30BD_u30FC_u30B9_u30B3_u30FC_u30C9_u3092_u78BA_u304B_u3081_u308B
/*
ポインタと値のレシーバに関するルールは、値のメソッドはポインタと値で呼び出すことができますが、
ポインタのメソッドはポインタにしか呼び出すことができません。

このルールは、ポインタメソッドはレシーバを変更することができるために発生します。
値に対してポインタメソッドを呼び出すと、メソッドは値のコピーを受け取ることになり、変更は破棄されてしまいます。
そのため、言語ではこのようなミスを禁止しています。しかし、便利な例外があります。
値がアドレス可能な場合、言語はアドレス演算子を自動的に挿入することで、ポインタメソッドを値上で呼び出すというよくあるケースに対処します。
この例では、変数bはアドレス指定可能なので、b.Writeだけで変数のWriteメソッドを呼び出すことができます。
コンパイラはこれを(&b).Writeに書き換えてくれます。
*/

type Human struct {
	Name string
}

func (h Human) Get() string {
	return h.Name
}

func (h Human) Set(name string) {
	fmt.Printf("value method: %p\n", &h)
	h.Name = name
}

func (h *Human) SetWithPointer(name string) {
	fmt.Printf("pointer method: %p\n", h)
	h.Name = name
}

func main() {
	//値型変数
	a := Human{Name: "init"}
	fmt.Println("値型変数の検証")

	fmt.Printf("before set call aのアドレス:%p\n", &a)
	a.Set("Hello1")
	fmt.Printf("name: %s\n", a.Get())

	fmt.Printf("before setWithPointer call aのアドレス:%p\n", &a)
	(&a).SetWithPointer("Hello2")
	fmt.Printf("name: %s\n", a.Get())

	fmt.Printf("before set call aのアドレス:%p\n", &a)
	a.SetWithPointer("Hello3")
	fmt.Printf("name: %s\n\n", a.Get())

	// ポインタ型変数
	b := &Human{Name: "init 2"}
	fmt.Println("ポインタ型変数の検証")

	fmt.Printf("before set call :%p\n", b)
	b.Set("Hello1")
	fmt.Printf("name: %s\n", (*b).Get())

	fmt.Printf("before setWithPointer call bが持つアドレス:%p\n", b)
	(*b).SetWithPointer("Hello2")
	fmt.Printf("name: %s\n", b.Get())

	fmt.Printf("before setWithPointer call bが持つアドレス:%p\n", b)
	b.SetWithPointer("Hello3")
	fmt.Printf("name: %s\n", b.Get())
}
