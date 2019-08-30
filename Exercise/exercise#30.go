package main

import "fmt"

func main() {
	f1 := func() int {
		return 100
	}
	str := foo(f1)

	fmt.Println("value returned from foo ", str)
}

func foo(f func() int) int {
	fn := f()
	fn++
	return fn
}
