package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Sample struct {
	A string `json:"a,omitempty"`
	B int    `json:"b"`
}

func main() {
	samples := []Sample{{"1", 1}, {B: 2}}
	f, _ := os.Create("sample.json")
	defer f.Close()

	err := json.NewEncoder(f).Encode(samples)
	if err != nil {
		log.Fatal(err)
	}

	f1, _ := os.Open("sample.json")
	defer f1.Close()
	var samples1 []Sample
	if err := json.NewDecoder(f1).Decode(&samples1); err != nil {
		log.Fatal(err)
	}
	fmt.Println(samples1)
}
