package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "net/http/pprof"
)

// 起動時にgo tool pprof leak http://localhost:6060/debug/pprof/heapで結果データを確認する
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("start")
	exec()
	fmt.Println("end")
}

func exec() {
	var someMap = make(map[string]string)

	addMap(someMap)
}

// すぐに1GBぐらいのメモリ食うので要注意
func addMap(s map[string]string) {
	i := int64(0)
	for {
		key := "key" + strconv.FormatInt(i, 10)
		val := "value" + strconv.FormatInt(i, 10)
		s[key] = val
		i++
	}
}
