package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./cicd/cache/heavy/greet/sample_gen.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("package greet\n\n")
	file.WriteString("import \"fmt\"\n\n")
	for i := 0; i < 10000; i++ {
		file.WriteString(fmt.Sprintf("func fn%d() {\n", i))
		file.WriteString(fmt.Sprintf("\tfmt.Println(\"Hello, World%d!\")\n", i))
		file.WriteString("}\n")
	}

	file.WriteString("\nfunc CallGreet() {\n")
	for i := 0; i < 10000; i++ {
		file.WriteString(fmt.Sprintf("\tfn%d()\n", i))
	}
	file.WriteString("}\n")
}
