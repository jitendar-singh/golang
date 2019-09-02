package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	fmt.Println("Version\t", runtime.Version())
	// fmt.Println("Memory Profile\t", runtime.ReadMemStats(&runtime.MemStats))
	fmt.Println("------------------------------")
	// foo()
	// bar()

	wg.Add(1)
	go foo()
	fmt.Println("Waiting List Go Routines ", runtime.NumGoroutine())
	wg.Wait()
	wg.Add(1)
	go bar()
	fmt.Println("Waiting List Go Routines ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Waiting List Go Routines ", runtime.NumGoroutine())
	fmt.Println("We are Done!!:)")

}

func foo() {
	fmt.Println("I am foo")
	wg.Done()

}

func bar() {
	fmt.Println("I am bar")
	wg.Done()

}
