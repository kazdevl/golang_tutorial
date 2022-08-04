package main

import (
	"testing"
)

func BenchmarkSample_GetDF(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ { // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
		s := Sample{
			A: 1,
			B: 1,
			C: false,
			D: "sample",
			E: true,
			F: "samples",
		}
		s.GetDF()
	}
}

func BenchmarkOptimizedSample_GetDF(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ { // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
		os := OptimizedSample{
			A: 1,
			B: 1,
			C: false,
			D: "sample",
			E: true,
			F: "samples",
		}
		os.GetDF()
	}
}

func Benchmark_GetSamples(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ { // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
		GetSamples(10000)
	}
}

func Benchmark_GetOptimizedSamples(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ { // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
		GetOptimizedSamples(10000)
	}
}
