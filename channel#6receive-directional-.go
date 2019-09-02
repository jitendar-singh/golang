package main

import "fmt"

func main() {

	c := make(chan int)
	// This is a receive only channel
	cr := make(<-chan int)
	// This is a send only channel
	cs := make(chan<- int)

	fmt.Printf("c\t%T\t\n", c)
	fmt.Printf("cs\t%T\t\n", cs)
	fmt.Printf("cr\t%T\t\n", cr)
}
