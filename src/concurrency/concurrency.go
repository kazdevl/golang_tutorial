package concurrency

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

// Goのランタイムに管理される軽量なスレッド
// goroutineは同じアドレス空間で実行されるので共有メモリへのアクセスは必ず同期させる必要がある
func GoRoutines() {
	oneThread := func(value string) {
		for i := 0; i < 3; i++ {
			fmt.Println(value)
			time.Sleep(3 * time.Second)
		}
	}
	go oneThread("goroutineを使って実行")
	oneThread("通常のスレッドで実行")
	fmt.Println("done")
}

// チャネル( Channel )型は、チャネルオペレータの <- を用いて値の送受信ができる通り道です。
// データは矢印の方向に流れる
// 通常、片方が準備できるまで送受信はブロックされます。これにより、明確なロックや条件変数がなくても、goroutineの同期を可能にします。
// バッファが詰まった時は、チャネルへの送信をブロックします。 バッファが空の時には、チャネルの受信をブロックします。
func Channels() {
	bufferNum := 3
	messages := make(chan string, bufferNum)
	go func() { messages <- "Hello" }()
	go func() { messages <- "World" }()
	messages <- "Sample"

	for i := 0; i < bufferNum; i++ {
		fmt.Println(<-messages)
	}
}

func ChannelsForDecentralize() {
	sum := func(s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum
	}
	c := make(chan int)
	data := []int{1, 2, 3, 4, 5, 6}
	go sum(data[:len(data)/2], c)
	go sum(data[len(data)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}

// 送り手は、これ以上の送信する値がないことを示すため、チャネルを close できます。
// 受け手は、受信の式に2つ目のパラメータを割り当てることで、そのチャネルがcloseされているかどうかを確認できます
// v, ok := <-ch
// 受信する値がない、かつ、チャネルが閉じているなら、 ok の変数は、 false になります。
// ※送り手のチャネルだけをcloseしてください。受け手はcloseしてはいけません。もしcloseしたチャネルへ送信すると、パニック( panic )します。
// ※チャネルはファイルとは異なり、通常はcloseする必要はありません。closeするのはこれ以上値が来ないことを受け手が知る必要があるときにだけです。例えばrange ループを終了するという場合です。
func RangeAndClose() {
	fn := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}

	c := make(chan int, 10)
	fmt.Printf("関数実行前の容量: %d, 長さ: %d\n", cap(c), len(c))
	go fn(cap(c), c)
	fmt.Printf("関数実行後の容量: %d, 長さ: %d\n", cap(c), len(c))
	for i := range c {
		fmt.Println(i)
	}
}

// select ステートメントは、goroutine(スレッド)を複数の通信操作で待たせます。
// select は、複数ある case のいずれかが準備できるようになるまでブロックし、準備ができた case を実行します。 もし、複数の case の準備ができている場合、 case はランダムに選択されます
// 参考文献: https://qiita.com/najeira/items/71a0bcd079c9066347b4,
func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	oneThread := func(sleeptime int, msg string, sending chan<- string) {
		time.Sleep(time.Duration(sleeptime) * time.Second)
		sending <- msg
	}

	go oneThread(5, "five second sleep", c2)
	go oneThread(3, "three second sleep", c1)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	fibonachi := func(c, quit chan int) {
		x, y := 0, 1
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
			default:
				fmt.Println("c and quit are blocked")
				// quit <- 0
			}
		}
	}
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonachi(c, quit)
}

// 6
// 参考文献: https://blog.toshimaru.net/goroutine-with-waitgroup/#goroutine--errgroup--context-%E3%82%92%E4%BD%BF%E3%81%86
func longTimeProcess(number int, reporter chan<- int) {
	time.Sleep(time.Second * 2)
	reporter <- number
}

func DetectFinishWithChannel() {
	reporter := make(chan int, 100)
	for i := 0; i < 100; i++ {
		go longTimeProcess(i, reporter)
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("END: %d\n", <-reporter)
	}
	fmt.Println("Finish Operation")

}

func ErrorHandlingByErrGroup() {
	var eg errgroup.Group
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(
			func() error {
				time.Sleep(2 * time.Second)
				if i > 90 {
					fmt.Println("Error:", i)
					return fmt.Errorf("error occured in %d", i)
				}
				fmt.Println("End:", i)
				return nil
			},
		)
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}

func ErrorHandlingByErrGroupAndContext() { //エラーが発生した時後続のgoroutineも削除する
	eg, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(
			func() error {
				time.Sleep(2 * time.Second)

				select {
				case <-ctx.Done():
					fmt.Println("Candeled:", i)
					return nil
				default:
					if i > 90 {
						fmt.Println("Error:", i)
						return fmt.Errorf("error occured in %d", i)
					}
					fmt.Println("End:", i)
					return nil
				}
			},
		)
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
