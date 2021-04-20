package basics

import (
	"fmt"
)

// 複数の引数で型が同様の場合、省略可能
func Add(x, y int) int {
	return x + y
}

// 複数の引数を返すことが可能
func Swap(a, b string) (string, string) {
	return b, a
}

// 複数の変数の最後に型を書くことで、変数のリスト宣言できる
func Variables() {
	var c, python, java bool
	var i int
	// 変数が宣言されて、初期化されてない場合, ゼロ値が入れられる
	// 数値型...int, float64: 0
	// bool型: false
	// string型: ""(empty string)
	fmt.Println(i, c, python, java)
}

// var 宣言では、変数毎に初期化子( initializer )を与えることができます。
// 初期化子が与えられている場合、型を省略できます。その変数は初期化子が持つ型になります。
func VariablesWithInitializers() {
	var c, python, java = true, false, "no!"
	var i int = 1
	fmt.Println(i, c, python, java)
}

// 関数の中では、 var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができます。
// 明示的な型を指定せずに変数を宣言する場合( := や var = のいずれか)、変数の型は右側の変数から型推論されます。
// なお、関数の外では、キーワードではじまる宣言( var, func, など)が必要で、 := での暗黙的な宣言は利用できません。
func ShortVariableDeclarations() {
	k := 3
	l := "少年"
	c, python, java := true, false, "no!"
	fmt.Println(k, l, c, python, java)
}

// 型変換
func TypeConversions() {
	var x, y int = 3, 4
	var convertX float64 = float64(x)
	var convertY float64 = float64(y)
	var convertUintX uint = uint(x)
	fmt.Println(convertX, convertY, convertUintX)
}

// 定数
// 定数( constant )は、 const キーワードを使って変数と同じように宣言します。
// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使えます。
// なお、定数は := を使って宣言できません。
func Constants() {
	const Name = "John"
	fmt.Println("Hello, ", Name)
}
