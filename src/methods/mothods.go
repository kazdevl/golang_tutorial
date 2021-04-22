package methods

import (
	"fmt"
	"math"
)

// methods
// classの仕組みはないが、型にメソッドを定義できる
// ポインタレシーバを使う2つの理由があります。
// ひとつは、メソッドがレシーバが指す先の変数を変更するためです。
// ふたつに、メソッドの呼び出し毎に変数のコピーを避けるためです。 例えば、レシーバが大きな構造体である場合に効率的です。
func Methods() {
	john := new(human)
	// レシーバー引数にポインタ型を代入したメソッド...レシーバー引数に代入した値自体を編集する場合は参照型にする
	john.setNameWithPointer("John") // 左は利便性のため、(*john).setNameWithPointerとして解釈される
	fmt.Println(john.getName())
	// レシーバー引数に値型を代入したメソッド...引数の型が値型であるために、引数に代入した値のコピーが生成される
	john.setName("Aloha")
	fmt.Println(john.getName())
}

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

// interfaces
// interface型は、メソッドのシグニチャの集まりで定義される
// そのメソッドの集まりを実装した値は、そのメソッドの集まりを持ったinterface型として振る舞える...ダックタイピング
// 参考文献：https://qiita.com/tenntenn/items/eac962a49c56b2b15ee8#%E5%9F%BA%E6%9C%AC%E7%9A%84%E3%81%AA%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%95%E3%82%A7%E3%83%BC%E3%82%B9%E3%81%AE%E5%AE%9F%E8%A3%85
func InterfaceImplementImplicity() {
	man := NewPerson("John", "Nick", isMale, 30)
	fmt.Printf("型の確認: %T\n", man)
	fmt.Println(man.GetAge(), man.GetName())
	woman := NewPerson("Jecy", "Smis", isFemale, 30)
	fmt.Printf("型の確認: %T\n", woman)
	fmt.Println(woman.GetAge(), woman.GetName())
}

func NewPerson(firstName, lastName string, gender Gender, age int) PersonI {
	name := &Name{FirstName: firstName, LastName: lastName}
	if gender == isFemale {
		return &Female{name, age}
	}
	return &Male{name, age}
}

type PersonI interface {
	GetName() string
	GetAge() int
}
type Name struct {
	FirstName string
	LastName  string
}

func (n *Name) GetName() string {
	return n.FirstName + " " + n.LastName
}

type Gender int

const (
	isFemale = iota
	isMale   = 1
)

type Male struct {
	*Name
	Age int
}

type Female struct {
	*Name
	Age int
}

func (male *Male) GetAge() int {
	return male.Age
}

func (female *Female) GetAge() int {
	return female.Age
}

// interface
// そのメソッドの集まりを実装した値を、interface型の変数へ持たせることができます。
// インタフェースの値は、特定の基底になる具体的な型の値を保持します。
func Interface() {
	var a Abser
	fmt.Printf("aの型: %T\n", a) //現時点では、aには何も値が代入されていないので、明確な型が存在しない
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Printf("aの型: %T, aの値: %v, abs: %F\n", a, a, a.Abs())
	a = &v // a *Vertext implements Abser
	fmt.Printf("aの型: %T, aの値: %v, abs: %F\n", a, a, a.Abs())
	// a = v v is aVertext (not *Vertex), and dosen't implemtn Abser
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f) //MyFloatはfloat64のように扱えるが、別の型として判断されるので、型キャストが必須
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// インターフェース自体の中にある具体的な値が nil の場合、メソッドは nil をレシーバーとして呼び出されます。
// いくつかの言語ではこれは null ポインター例外を引き起こしますが、Go では nil をレシーバーとして呼び出されても適切に処理するメソッドを記述するのが一般的です(この例では M メソッドのように)。
// 具体的な値として nil を保持するインターフェイスの値それ自体は非 nil であることに注意してください。
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func InterfaceValuesWithNil() {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	// varで変数を定義するだけでは、値は生成されない
	// newで値が生成され、ポインタ型の場合はそれへのアドレスが代入される。
	newT := new(T)
	describe(newT)
	newT.M()
	fmt.Printf("tが持つアドレス: %p, t自身のアドレス; %p\nnewTが持つアドレス: %p, newT自身のアドレス: %p\n", t, &t, newT, &newT)

	i = &T{"HELLO"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//  nil インターフェースの値は、値も具体的な型も保持しません。
// 呼び出す 具体的な メソッドを示す型がインターフェースのタプル内に存在しないため、
// nil インターフェースのメソッドを呼び出すと、ランタイムエラーになります。
func NilInterfaceValues() {
	var i I
	describe(i)
	i.M()
}

// type assertions
// インタフェースの値の元になる具体的な値を利用する手段を提供している
func TypeAssertions() {
	//0個のメソッドを指定されたインタフェース型は、空のインタフェースと呼ばれいる
	// 空のインタフェースは、任意の型の値を保持できる...全ての型は、少なくとも0個のメソッドを実装しているため
	// 未知の値の型を取り扱うときに利用する
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	// f = i.(float64): panic: interface conversion: interface {} is string, not float64というエラーが発生する
}

func TypeSwitches() {
	checkType("Hello")
	checkType(100)
	checkType(true)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("this type is int. twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}
