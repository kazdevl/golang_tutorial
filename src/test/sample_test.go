package sample_test

import (
	"fmt"
	"testing"
)

// 参考文献1: https://golang.org/pkg/testing/#B
// 参考文献2: https://golang.org/cmd/go/#hdr-Testing_flags

func TestFalsePattern(t *testing.T) {
	expectV := 1
	actualV := 2
	if expectV != actualV {
		t.Errorf("The actual value is %d, but the expected value is %d", actualV, expectV)
	}
}
func TestSuccessPattern(t *testing.T) {
	expectV := 1
	actualV := 1
	if expectV != actualV {
		t.Errorf("The actual value is %d, but the expected value is %d", actualV, expectV)
	}
}

func BenchmarkAppend_AllocateEveryTime(b *testing.B) {
	base := []string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ { // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
		base = append(base, fmt.Sprintf("No %d", i))
	}
}

func BenchmarkAppend_AllocateOnce(b *testing.B) {
	base := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base[i] = fmt.Sprintf("No %d", i)
	}
}
