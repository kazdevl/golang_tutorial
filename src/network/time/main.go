package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go launchMockServer()
	go timeOut(ctx)
	go calcTime(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("time out")
	}
}

func launchMockServer() {
	http.HandleFunc("/sleep/threesecond", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "hello world"}`))
	})
	http.ListenAndServe(":8080", nil)
}

func timeOut(ctx context.Context) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/sleep/threesecond", nil)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	msg, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("message in timeOut: %s\n", msg)
}

func calcTime(ctx context.Context) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/sleep/threesecond", nil)
	client := &http.Client{}
	start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	diff := time.Since(start).Milliseconds()
	fmt.Printf("RTT: %v(ms)\n", diff)
	defer res.Body.Close()
	msg, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("message in calcTime: %s\n", msg)
}
