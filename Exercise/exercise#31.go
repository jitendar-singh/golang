package main

import "fmt"

func main() {

	var x *int
	var a int
	x = &a
	*x++

	fmt.Println("using a ", a)
	fmt.Println("using x ", *x)
}
