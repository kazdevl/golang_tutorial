package fileRead_test

import (
	"app/kernel/fileRead"
	"encoding/json"
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

	json1, err := os.Create("sample1.json")
	defer json1.Close()
	if err != nil {
		log.Fatal(err)
	}
	cs := make([]fileRead.Content, 1000)
	for i := 0; i < 1000; i++ {
		cs[i].Page = fmt.Sprintf("http://localhost:8080/%d", i)
	}
	if err := json.NewEncoder(json1).Encode(&cs); err != nil {
		log.Fatal(err)
	}

	json2, err := os.Create("sample2.json")
	defer json2.Close()
	if err != nil {
		log.Fatal(err)
	}
	scs := fileRead.SliceContent{}
	scs.Pages = make([]string, 1000)
	for i := 0; i < 1000; i++ {
		scs.Pages[i] = fmt.Sprintf("http://localhost:8080/%d", i)
	}
	if err := json.NewEncoder(json2).Encode(&scs); err != nil {
		log.Fatal(err)
	}
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

func Benchmark_ReadFile(b *testing.B) {
	targets := []struct {
		name string
		fn   func() []string
	}{
		{
			name: "text file read",
			fn:   fileRead.ReadTextFileContent,
		},
		{
			name: "json file read with contents",
			fn:   fileRead.ReadJsonFileContentWithContents,
		},
		{
			name: "json file read with slice content",
			fn:   fileRead.ReadJsonFileContentWithSliceContent,
		},
	}

	for _, target := range targets {
		b.Run(target.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				target.fn()
			}
		})
	}
	b.ResetTimer()
}
