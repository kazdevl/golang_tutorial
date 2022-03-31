package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants"
)

var globalSum int32

type Number interface {
	int8 | int16 | int32 | int64
}

func myFunc[N Number](i N) {
	atomic.AddInt32(&globalSum, int32(i))
	fmt.Printf("run with %v \n", i)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	runTimes := 1000

	// use common pool
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Print("finish all tasks.\n")

	// use the pool with a function,
	p, _ := ants.NewPoolWithFunc(10, func(i any) {
		addData := i.(int32)
		myFunc[int32](addData)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks. result is %d\n", globalSum)
}
