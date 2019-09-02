package main

import "fmt"

func main() {
	//This is a send only type channel
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)

}
