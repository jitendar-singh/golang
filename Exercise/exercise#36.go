package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("This is foo executing")
		fmt.Println("Number of Go Routines:", runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {
		fmt.Println("This is bar executing")
		fmt.Println("Number of Go Routines:", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())
	fmt.Println("We are done now!!")

}
