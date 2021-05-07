package effective_go

import (
	"fmt"
)

// 参考文献1: https://golang.org/doc/effective_go#methods
// 参考文献2: https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/#u30B3_u30F3_u30D1_u30A4_u30E9_u306E_u30BD_u30FC_u30B9_u30B3_u30FC_u30C9_u3092_u78BA_u304B_u3081_u308B
// 参考文献3: https://qiita.com/nirasan/items/02e88c3ba64c444fa527...値がアドレス可能でないパターンの把握(要確認)
// 参考文献4: https://qiita.com/tikidunpon/items/2d9598f33817a6e99860
// 参考文献5: https://github.com/golang/go/wiki/MethodSets#interfaces
/*
ポインタと値のレシーバに関するルールは、値のメソッドはポインタと値で呼び出すことができますが、
ポインタのメソッドはポインタにしか呼び出すことができません。

このルールは、ポインタメソッドはレシーバを変更することができるために発生します。
値に対してポインタメソッドを呼び出すと、メソッドは値のコピーを受け取ることになり、変更は破棄されてしまいます。
そのため、言語ではこのようなミスを禁止しています。しかし、便利な例外があります。
値がアドレス可能な場合、言語はアドレス演算子を自動的に挿入することで、ポインタメソッドを値上で呼び出すというよくあるケースに対処します。...値がアドレス可能でない時とはいつなのかのチェック(値型でマップやインターフェイスの要素)
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

func (h *Human) CallForNil(name string) string {
	fmt.Printf("pointer method: %p\n", h)
	return "Hi! nil, my name is " + name
}

func Check() {
	//値型変数
	a := Human{Name: "init"}
	fmt.Println("値型変数の検証")
	// 値型メソッドがレシバーの値をコピーしているかの確認
	fmt.Printf("before set call aのアドレス:%p\n", &a)
	a.Set("Hello1")
	fmt.Printf("name: %s\n", a.Get())
	// 以下においては、値がアドレス可能な場合、値がコンパイラでアドレス演算子を自動的に挿入しているかの確認
	fmt.Printf("before setWithPointer call aのアドレス:%p\n", &a)
	(&a).SetWithPointer("Hello2")
	fmt.Printf("name: %s\n", a.Get())

	fmt.Printf("before set call aのアドレス:%p\n", &a)
	a.SetWithPointer("Hello3")
	fmt.Printf("name: %s\n\n", a.Get())

	// ポインタ型変数
	b := &Human{Name: "init 2"}
	fmt.Println("ポインタ型変数の検証")
	// 値のメソッドはポインタでもしっかりと呼び出されるのかの検証...暗黙的変換の検証
	fmt.Printf("before set call bが持つアドレス:%p\n", b)
	b.Set("Hello1")
	fmt.Printf("name: %s\n", (*b).Get())
	fmt.Printf("name: %s\n", b.Get())

	// 構造体へのポインタ変数の場合、参照先のフィールドにアクセスする場合、コンパイラがb.Nameを、(*b).Nameに変換してくれる
	fmt.Printf("before setWithPointer call bが持つアドレス:%p\n", b)
	b.SetWithPointer("Hello2")
	fmt.Printf("name: %s\n", b.Name)
	fmt.Printf("name: %s\n\n", (*b).Name)

	// nilの検証...ポインタ編
	var c *Human
	fmt.Println("nil...ポインタ型の検証")
	fmt.Printf("before CallForNil call cが持つアドレス:%p\n", c)
	fmt.Printf("name: %s\n\n", c.CallForNil("Test"))

	// nilの検証...値編
	var d Human
	fmt.Println("nil...値型の検証")
	fmt.Printf("before set call dのアドレス:%p\n", &d)
	fmt.Printf("name: %s\n\n", d.CallForNil("Test"))

	// mapとinterfaceの検証
	fmt.Println("mapとinterfaceの検証")
	// TODO
}


type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}
// インターフェイスに格納されている具象値は、マップ要素がアドレス指定できないのと同じように、アドレス指定できません。
// ポインタレシーバーのメソッドを値で呼び出せないのは、インターフェイス内に格納された値がアドレスを持たないためです。インターフェイスに値を代入する際、
// コンパイラは可能なすべてのインターフェイス・メソッドが実際にその値で呼び出されることを___保証__するため、不適切な代入を行おうとするとコンパイル時に失敗します。...interfaceは、interfaceが定義したメソッドを実装したtypeであれば問題なく代入でき、代入されるものが全てアドレスを持つのか判別できない
func Check2() {
	// A bare value
	var lst List = List{1, 2, 3}
	CountInto(lst, 1, 10) // INVALID: Append has a pointer receiver
	if LongEnough(lst) {  // VALID: Identical receiver type
		fmt.Printf(" - lst is long enough")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) // VALID: Identical receiver type
	if LongEnough(plst) {  // VALID: a *List can be dereferenced for the receiver
		fmt.Printf(" - plst is long enough")
	}
}
