package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	const gs = 1000000

	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("Total Counter:", counter)
	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())

}
