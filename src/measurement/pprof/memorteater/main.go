package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/profile"
)

// go build -gcflags '-m' main.goでヒープに保存されるか確認できる
// go tool pprof memory mem.pprofで解析できる
//  その後topで全体像をlistで行毎に置けるメモリの使用量を確認できる
func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	fmt.Println("start")
	exec()
	fmt.Println("end")
}

func exec() {
	var someMap = make(map[string]string)
	addMap(someMap, "hello", 100000)
	doNothing("world", 100000)
}

func addMap(s map[string]string, prefix string, count int) {
	for i := 0; i < count; i++ {
		key := prefix + "key" + strconv.Itoa(count)
		val := "value" + strconv.Itoa(i)
		s[key] = val
	}
}

func doNothing(prefix string, count int) {
	for i := 0; i < count; i++ {
		key := prefix + "key" + strconv.Itoa(count)
		val := "value" + strconv.Itoa(i)
		key = key + val
	}
}
