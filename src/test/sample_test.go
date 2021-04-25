package sample_test

import (
	"fmt"
	"testing"
)

// 参考文献1: https://golang.org/pkg/testing/#B
// 参考文献2: https://golang.org/cmd/go/#hdr-Testing_flags
// 参考文献3: https://future-architect.github.io/articles/20200601/

// tesingはassertを持っておらず、自分で生成したエラーを自ら返すことでtestを行っている
// assertで帰ってくる値は、開発者にとって欲しい情報でない場合が多く、自らカスタマイズするべきであると考えているため
// また、testは失敗した時の情報こそ大事と考えているので、成功パターン時に帰ってくる情報は多くはない。
// https://golang.org/doc/faq#assertions
func TestFalsePattern(t *testing.T) {
	expectV := 1
	actualV := 2
	if expectV != actualV {
		t.Errorf("The actual value is %d, but the expected value is %d", actualV, expectV)
	}
}
func TestSuccessPattern(t *testing.T) {
	expectV := 1
	actualV := 1
	if expectV != actualV {
		t.Errorf("The actual value is %d, but the expected value is %d", actualV, expectV)
	}
}

// The benchmark function must run the target code b.N times. During benchmark execution
// b.N is adjusted until the benchmark function lasts long enough to be timed reliably
// the output example
// benchmarkname loopRunTimes aSpeedPerLoop
func BenchmarkAppend_AllocateEveryTime(b *testing.B) {
	base := []string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ { // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
		base = append(base, fmt.Sprintf("No %d", i))
	}
}

func BenchmarkAppend_AllocateOnce(b *testing.B) {
	base := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base[i] = fmt.Sprintf("No %d", i)
	}
}

// Examle
// "Output"で始まるコメントを含んでおり、テストが実行された際の関数の標準出力と比較される
// outputコメントを含まないExample関数はコンパイルされるが実行されない
// パッケージ・関数・タイプ・タイプにあるメソッドの例を宣言するための命名規則がある
// 出力の結果をテストしたい場合に利用する or godocにExampleを表示させたい場合にも利用する
func Example() {
	fmt.Println("Hello")
	// Output: Hello
}

// Skip & Subtests
func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "success", args: args{a: 1, b: 2}, want: 3},
		{name: "fail", args: args{a: 1, b: 2}, want: 30},
	}

	for _, test := range tests {
		// tの名前のサブテストとして、fを実行する
		// fを別のgoroutineで実行し、fが戻る or t.Parallel()を呼び出して、並列testになるまでブロックする
		// runはfが成功したかを報告する
		t.Run(test.name, func(t *testing.T) {
			// testとそのサブテストが完了した時に、呼び出される関数を登録する
			// LIFO
			// t.FatalFでテストが失敗しても呼び出される
			t.Cleanup(func() {
				// テキストをエラーログに記録する
				// テストでは、testが失敗した場合や、-test.vのフラグが設定されている場合のみ、テキストが出力される
				// ベンチマークでは、-test.vフラグの値に、パフォーマンスが依存するのを避けるために、常にテキストが出力される
				t.Log("clean up!")
			})
			// t.FatalFでテストが失敗しても呼び出される
			defer t.Log("defer!")
			// testをskipする
			// Short()は-test.shortのフラグが設定されているか否かをboolで返す
			if testing.Short() {
				t.Skip("skipping test in short mode")
			}
			// test
			if result := add(test.args.a, test.args.b); result != test.want {
				t.Fatalf("add() = %d, but want is %d", result, test.want)
			}
			// t.Fatalfでテストが失敗した場合、以下は呼ばれない
			t.Log("t.Fatalf is not called")
		})
	}
}

// Subbenchmarks
func BenchmarkAddd_Allocate(b *testing.B) {
	targets := []struct {
		name string
		fn   func(int)
	}{
		{name: "allocate once", fn: func(loopNum int) {
			base := []string{}
			for i := 0; i < loopNum; i++ {
				base = append(base, fmt.Sprintf("No %d", i))
			}
		}},
		{name: "allocate everytime", fn: func(loopNum int) {
			base := make([]string, loopNum)
			for i := 0; i < cap(base); i++ {
				base[i] = fmt.Sprintf("No %d", i)
			}
		}},
	}

	for _, target := range targets {
		b.Run(target.name, func(b *testing.B) {
			b.ResetTimer()
			b.Cleanup(func() {
				b.Log("clean up!")
			})
			target.fn(b.N)
		})
	}
}
