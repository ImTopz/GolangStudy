package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func DoTask(wg *sync.WaitGroup) int {
	n := 2
	for i := 0; i < 20000; i++ {
		for j := 0; j < 100000; j++ {
			if n > 1000000 {
				n = n - 10000000
			} else {
				n++
			}
		}
	}
	(*wg).Done()
	return n
}
func DoTasks(x int) {
	runtime.GOMAXPROCS(x)
	var wg sync.WaitGroup
	start := time.Now().UnixNano()
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go DoTask(&wg)
	}

	wg.Wait()
	fmt.Println("cpu", x, time.Now().UnixNano()-start, "ns")
}
func main() {
	for i := 1; i <= 8; i++ {
		DoTasks(i)
	}
}
