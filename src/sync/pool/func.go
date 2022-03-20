package pool

import "sync"

var cap int = 1000
var GlobalL = make([]int, 0, cap)
var GlobalMutex = sync.Mutex{}
var GlobalPool = sync.Pool{
	New: func() any {
		return make([]int, 0, cap)
	},
}

// ref: https://christina04.hatenablog.com/entry/2016/12/29/224655

func CreateListEach() {
	l := make([]int, 0, cap)
	for i := 0; i < 5*cap; i++ {
		l = append(l, i)
	}
}

func CreateListOnce() {
	// can not use with concurrency
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()
	for i := 0; i < 5*cap; i++ {
		GlobalL = append(GlobalL, i)
	}

	GlobalL = GlobalL[:0]
}

func UseSyncPool() {
	l := GlobalPool.Get().([]int)
	for i := 0; i < 5*cap; i++ {
		l = append(l, i)
	}

	l = l[:0]
	GlobalPool.Put(l)
}
