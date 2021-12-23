package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			i += 1
			for j := 0; j < i; j++ {
				req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
				c := &http.Client{}
				c.Do(req)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go func() {
		for {
			cmd := exec.Command("docker", "stats", "--no-stream", "some-nginx")
			a, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			contents := strings.Split(string(a), "\n")
			cs := strings.Split(contents[1], " ")
			data := cs[6][:len(cs[6])-1]
			fmt.Println(data)
			// dataF, err := strconv.ParseFloat(data, 64)
			time.Sleep(5 * time.Second)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

// TODO 初期化についての調査
