package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"time"
	"log"
	"io/ioutil"
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

	cmd := exec.Command("docker", "stats", "--format", `"{{.Name}} {{.CPUPerc}} {{.NetIO}}"`, "some-nginx", "some-nginx-1")
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	go func () {
		for {
			time.Sleep(1*time.Second)
			b, _ := ioutil.ReadAll(pipe)
			fmt.Println(string(b))
			// conStents := strings.Split(string(b), "\n")
			// for _, content := range contents[:len(contents)-1] {
			// 	fmt.Println("****************")
			// 	cs := strings.Split(content, " ")
			// 	for _, c := range cs {
			// 		fmt.Println("--------")
			// 		fmt.Println(c)
			// 	}
		
			// 	// data := cs[1][:len(cs[1])-1]
			// 	// fmt.Println(data)
			// }
		}
	}()

	if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
