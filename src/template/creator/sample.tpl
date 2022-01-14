package main

import "fmt"

func main() {
    {{range .}}
    const {{.Name}} = "{{.Value}}"
    fmt.Println({{.Name}})
    {{end}}
}