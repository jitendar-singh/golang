package main

import "fmt"

func main() {
	var x func()
	x = foo()
	fmt.Println("Executing the varibale")
	x()
}

func foo() func() {
	fmt.Println("I am inside foo ")
	fmt.Println("Returning function bar")

	return func() {
		fmt.Println("Hello I am inside the returned func")
	}
}
