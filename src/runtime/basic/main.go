package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	fmt.Println("runtime package")
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Second().Do(func() {
		fmt.Println("1s")
		fmt.Printf("num fo goroutine: %d\n", runtime.NumGoroutine())
	})
	s.Every(3).Second().Do(func() {
		fmt.Println("3s")
		fmt.Printf("num fo goroutine: %d\n", runtime.NumGoroutine())
	})
	s.Every(5).Second().Do(func() {
		fmt.Println("5s")
		fmt.Printf("num fo goroutine: %d\n", runtime.NumGoroutine())
	})
	s.StartAsync()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			go func() {
				for {
					time.Sleep(2 * time.Second)
					fmt.Println("Hey!")
				}
			}()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("finish")
}
