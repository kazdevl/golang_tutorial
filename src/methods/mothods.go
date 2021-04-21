package methods

import "fmt"

type human struct {
	id   int
	age  int
	name string
}

func (h *human) getName() string {
	return h.name
}

func (h *human) setNameWithPointer(name string) {
	h.name = name
}

func (h human) setName(name string) {
	h.name = name
}

// methods
// classの仕組みはないが、型にメソッドを定義できる
func Methods() {
	john := new(human)
	// レシーバー引数にポインタ型を代入したメソッド...レシーバー引数に代入した値自体を編集する場合は参照型にする
	john.setNameWithPointer("John")
	fmt.Println(john.getName())
	// レシーバー引数に値型を代入したメソッド...引数の型が値型であるために、引数に代入した値のコピーが生成される
	john.setName("Aloha")
	fmt.Println(john.getName())
}
