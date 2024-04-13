package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

type Target struct {
	Name     string
	Kind     string `faker:"fixed"`
	Age      int    `faker:"boundary_start=30, boundary_end=50"`
	Features []feature
}

type feature struct {
	Value string
}

func CreateCutomGenerator() {
	faker.AddProvider("fixed", func(v reflect.Value) (interface{}, error) {
		return "sample", nil
	})
}

func main() {
	CreateCutomGenerator()
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
