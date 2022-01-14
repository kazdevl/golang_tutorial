package main

import "fmt"

func main() {
    {{range .}}
    const {{.Name}} = "{{.Value}}"
    if {{.Value}} % 2 == 0 {
        fmt.Println("偶数: ", {{.Name}})
    }
    fmt.Println("奇数: ", {{.Name}})
    {{end}}
}