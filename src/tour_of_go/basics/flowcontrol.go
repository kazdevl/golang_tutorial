package basics

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// for文
// 初期化(任意) + 条件式(任意...無限ループになる) + 後処理(任意)
// while的なことをfor文でできる
func For() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	// for {} //無限ループ
	fmt.Println(sum)
}

// if
// if ステートメントは、 for のように、条件の前に、評価するための簡単なステートメントを書くことができます。
// ここで宣言された変数は、 if のスコープ内だけで有効です。
func IfWithAShortStatement(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// switch
// 選択されたcaseのみ実行し、breakステートメントが自動で提供されてる
// caseは定数である必要ない
func Switch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux .")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// defer
// defer ステートメントは、 defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるものです。
// defer へ渡した関数の引数は、すぐに評価されますが、その関数自体は呼び出し元の関数がreturnするまで実行されません。
// defer へ渡した関数が複数ある場合、その呼び出しはスタック( stack )されます。 呼び出し元の関数がreturnするとき、 defer へ渡した関数は LIFO(last-in-first-out) の順番で実行されます。
func Defer() {
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
