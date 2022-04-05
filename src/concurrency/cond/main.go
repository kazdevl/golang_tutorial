package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	l := new(sync.Mutex)
	c := sync.NewCond(l)

	for i := 0; i < 10; i++ {
		go func(num int) {
			output := fmt.Sprintf("go routine %d", num)
			fmt.Printf("raise %s\n", output)
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fmt.Printf("finish %s\n", output)
		}(i)
	}

	fmt.Println("wait 5 second")
	time.Sleep(5 * time.Second)

	// for i := 0; i < 10; i++ {
	// 	c.Signal()
	// 	time.Sleep(500 * time.Millisecond)
	// }
	c.Broadcast()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("finish main go routine")
}
