package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	sub := "alice"
	obj := "data1"
	act := "read"

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Fatal(err)
	}

	if ok {
		fmt.Println("権限があります")
	} else {
		fmt.Println("権限がありません")
	}
}
