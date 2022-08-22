package main

import (
	"fmt"
	"unsafe"
)

type Sample struct {
	A int32
	B int64
	C bool
	D string
	E bool
	F string
}

type OptimizedSample struct {
	B int64
	D string
	F string
	A int32
	C bool
	E bool
}

func main() {
	s := Sample{}
	fmt.Printf(`
Sample Struct=%d(byte)
************************
A=%d(byte)
B=%d(byte)
C=%d(byte)
D=%d(byte)
E=%d(byte)
F=%d(byte)
Max Fields Byte=%d(byte)
`,
		unsafe.Sizeof(s),
		unsafe.Sizeof(s.A),
		unsafe.Sizeof(s.B),
		unsafe.Sizeof(s.C),
		unsafe.Sizeof(s.D),
		unsafe.Sizeof(s.E),
		unsafe.Sizeof(s.F),
		unsafe.Sizeof(s.A)+unsafe.Sizeof(s.B)+unsafe.Sizeof(s.C)+unsafe.Sizeof(s.D)+unsafe.Sizeof(s.E)+unsafe.Sizeof(s.F),
	)
	os := OptimizedSample{}
	fmt.Println("----------------")
	fmt.Printf(`
Optimized Sample Struct=%d(byte)
************************
A=%d(byte)
B=%d(byte)
C=%d(byte)
D=%d(byte)
E=%d(byte)
F=%d(byte)
Max Fields Byte=%d(byte)
			`,
		unsafe.Sizeof(os),
		unsafe.Sizeof(os.A),
		unsafe.Sizeof(os.B),
		unsafe.Sizeof(os.C),
		unsafe.Sizeof(os.D),
		unsafe.Sizeof(os.E),
		unsafe.Sizeof(os.F),
		unsafe.Sizeof(os.A)+unsafe.Sizeof(os.B)+unsafe.Sizeof(os.C)+unsafe.Sizeof(os.D)+unsafe.Sizeof(os.E)+unsafe.Sizeof(os.F),
	)
}

func (s Sample) GetDF() string {
	return s.D + s.F
}

func (os OptimizedSample) GetDF() string {
	return os.D + os.F
}

func GetSamples(num int) []Sample {
	ss := make([]Sample, 0, num)
	for i := 0; i < num; i++ {
		ss = append(ss, Sample{
			A: 1,
			B: 1,
			C: false,
			D: "sample",
			E: true,
			F: "samples",
		})
	}
	return ss
}

func GetOptimizedSamples(num int) []OptimizedSample {
	oss := make([]OptimizedSample, 0, num)
	for i := 0; i < num; i++ {
		oss = append(oss, OptimizedSample{
			A: 1,
			B: 1,
			C: false,
			D: "sample",
			E: true,
			F: "samples",
		})
	}
	return oss
}
