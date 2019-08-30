package main

import "fmt"

func main() {

	var f func()
	f = hello
	f()

}

func hello() {
	fmt.Println("Hello form f")
}
