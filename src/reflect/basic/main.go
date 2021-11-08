package main

import (
	"fmt"
	"reflect"
)

// ref: https://www.slideshare.net/takuyaueda967/reflect-27186813
//      https://go.dev/blog/laws-of-reflection
func main() {
	type sample struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	var (
		a string  = "sample1"
		b sample  = sample{A: "sample2"}
		c *sample = &sample{A: "sample3"}
	)
	fmt.Println(typeAndKind(a))
	fmt.Println(valueAndKind(a))
	fmt.Println("************")
	fmt.Println(typeAndKind(b))
	fmt.Println(valueAndKind(b))
	fmt.Println("************")
	fmt.Println(typeAndKind(c))
	fmt.Println(valueAndKind(c))
	fmt.Println("************")
	d := 100
	v1 := reflect.ValueOf(d)
	fmt.Println(v1.CanSet())
	v2 := reflect.ValueOf(&d)
	v2.Elem().SetInt(200)
	fmt.Println(d)
	type sampleSample struct {
		C      string `json:"c_json"`
		Sample sample `json:"sample"`
	}
	e := sampleSample{C: "sample5", Sample: sample{A: "sample6", B: 5}}
	v3 := reflect.ValueOf(e)
	fmt.Println(v3.Field(0))
	fmt.Println(v3.FieldByIndex([]int{1, 0}))
	fmt.Println(v3.FieldByIndex([]int{1, 1}))
	fmt.Println(v3.FieldByName("C"))
	fmt.Println(v3.FieldByName("Sample"))
	fmt.Println(e)
	v4 := reflect.ValueOf(&e)
	fmt.Println(v4.Elem().Field(0))
	v4.Elem().FieldByName("C").SetString("HogeHoge")
	v4.Elem().FieldByName("Sample").Field(0).SetString("Hoge")
	v4.Elem().FieldByName("Sample").Field(1).SetInt(100)
	fmt.Println(e)
	v5 := reflect.TypeOf(e)
	n, _ := v5.FieldByName("C")
	fmt.Println(n.Tag.Lookup("json"))
	fmt.Println(n.Tag.Get("json"))
	fmt.Println(n.Tag.Lookup("hoge"))
	fmt.Println(n.Tag.Get("hoge"))
}

func typeAndKind(v interface{}) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(v)
	k := t.Kind()

	return t, k
}

func valueAndKind(v interface{}) (reflect.Value, reflect.Kind) {
	t := reflect.ValueOf(v)
	k := t.Kind()
	return t, k
}
