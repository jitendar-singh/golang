package main

import "fmt"

func main() {

	//This allocates a channel & blocks buffer for 2 values

	c := make(chan string, 2)

	c <- "Jitendar"
	c <- "Singh"

	fmt.Println(<-c, <-c)
}
