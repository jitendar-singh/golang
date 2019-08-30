package main

import "fmt"

func main() {
	fmt.Println(foo(6))
	fmt.Println(bar(5, "Jitendar"))
}

func foo(x int) int {
	return x
}

func bar(x int, name string) (int, string) {
	return x, name
}
