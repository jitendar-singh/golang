package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU Profile ", runtime.NumCPU())
	fmt.Println("Go Routines ", runtime.NumGoroutine())

	counter := 0

	const gs = 10000000

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routines ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routines ", runtime.NumGoroutine())
	fmt.Println("Count:", counter)

}
