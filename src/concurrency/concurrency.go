package concurrency

import (
	"fmt"
	"time"
)

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

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	oneThread := func(sleeptime int, msg string, sending chan<- string) {
		time.Sleep(2 * time.Second)
		sending <- msg
	}

	go oneThread(2, "two second sleep", c2)
	go oneThread(1, "one second sleep", c1)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
