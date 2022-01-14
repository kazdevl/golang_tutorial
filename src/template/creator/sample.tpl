package main

import "fmt"

func main() {
    {{range .}}
    const {{.Name}} = "{{.Value}}"
    const {{.Name}}_int = {{.Value}}
    if {{.Name}}_int % 2 == 0 {
        fmt.Println("偶数: ", {{.Name}})
    }
    fmt.Println("奇数: ", {{.Name}})
    {{end}}
}