package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 100; i++ {
		createFile(i)
	}

	f, err := os.Create("./cicd/cache/heavy/greet/all_gen.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("package greet\n\n")

	f.WriteString("func CallGreetAll() {\n")
	for i := 0; i < 100; i++ {
		f.WriteString(fmt.Sprintf("CallGreet%d()\n", i))
	}
	f.WriteString("}\n")
}

func createFile(id int) {
	file, err := os.Create(fmt.Sprintf("./cicd/cache/heavy/greet/sample_gen_%d.go", id))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("package greet\n\n")
	file.WriteString("import \"fmt\"\n\n")

	file.WriteString(fmt.Sprintf("\nfunc CallGreet%d() {\n", id))
	for i := 0; i < 10000; i++ {
		file.WriteString(fmt.Sprintf("fn_%d_%d()\n", id, i))
	}
	file.WriteString("}\n")

	for i := 0; i < 10000; i++ {
		file.WriteString(fmt.Sprintf("func fn_%d_%d() {\n", id, i))
		file.WriteString(fmt.Sprintf("\tfmt.Println(\"Hello, World%d!\")\n", i))
		file.WriteString("}\n")
	}
}
