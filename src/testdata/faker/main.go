package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

type Target struct {
	Name     string
	Age      int `faker:"boundary_start=30, boundary_end=50"`
	Features []feature
}

type feature struct {
	Value string
}

func main() {
	for range 3 {
		var t Target
		_ = faker.FakeData(&t, options.WithRandomMapAndSliceMinSize(1), options.WithRandomMapAndSliceMaxSize(3))
		m, _ := json.MarshalIndent(t, "", "  ")
		fmt.Printf("Target: %s\n", m)
	}

	fmt.Println("*****")
	vs := createDummys()
	for _, v := range vs {
		m, _ := json.MarshalIndent(v, "", "  ")
		fmt.Printf("WithTags: %s\n", m)
	}

	fmt.Println("*****")
	vs2 := createDummyUniques()
	for _, v := range vs2 {
		m, _ := json.MarshalIndent(v, "", "  ")
		fmt.Printf("Unique: %s\n", m)
	}

	fmt.Println("*****")
	v3 := createDummyLang()
	m, _ := json.MarshalIndent(v3, "", "  ")
	fmt.Printf("Lang: %s\n", m)

	for range 5 {
		fmt.Println("date: ", faker.Date())
	}
}
