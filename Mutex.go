package main

// Run -- go run -race Race.go

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
		
			// Gosched yields the processor, allowing other goroutines to run. 
			// It does not suspend the current goroutine, so execution resumes 
			// automatically.

			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)

}
