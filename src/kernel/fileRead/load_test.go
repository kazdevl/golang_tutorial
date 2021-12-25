package fileRead_test

import (
	"app/kernel/fileRead"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestFiles() {
	f, err := os.Create("sample1.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	strs := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		strs[i] = fmt.Sprintf("http://localhost:8080/%d", i)
	}
	f.WriteString(strings.Join(strs, "\n"))
}

func Test_ReadTextFileContent(t *testing.T) {
	createTestFiles()
	want := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		want[i] = fmt.Sprintf("http://localhost:8080/%d", i)
	}
	result := fileRead.ReadTextFileContent()
	assert.Equal(t, result, want)
}

func Benchmark_ReadTextFileContent(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fileRead.ReadTextFileContent()
	}
}
