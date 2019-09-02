package main

import "fmt"

func main() {

	//This allocates  a channel and blocks room for one value
	// c := make(chan string, 2) will block room for two values
	c := make(chan string, 1)

	c <- "Jitendar Singh"
	c <- "Sunny"
	fmt.Println(<-c)

}
