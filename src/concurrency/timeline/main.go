package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseContent struct {
	Addr string `json:"address"`
}

func main() {
	// launch server
	go launchMockServer()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	carIDs := []string{"c1", "c2", "c3", "c4", "c5"}
	for _, carID := range carIDs {
		carID := carID
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				execRequest(ctx, carID)
			}
		}()
	}
	<-ctx.Done()
	fmt.Println("time out")

}

func launchMockServer() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		longExec()
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"address": "sample"}`))
	}))
}

func longExec() {
	time.Sleep(200 * time.Millisecond)
}

func execRequest(ctx context.Context, carID string) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8080", nil)
	c := &http.Client{}
	start := time.Now()
	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("error request exec: %v\n", err)
		return
	}
	duration := time.Since(start).Milliseconds() //ms
	fmt.Printf("carID[%s]...time duration: %dms\n", carID, duration)
	defer res.Body.Close()
	var data ResponseContent
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Printf("error decode: %v\n", err)
		return
	}
	fmt.Printf("carID[%s]...parsed content: %+v\n", carID, data)
}
