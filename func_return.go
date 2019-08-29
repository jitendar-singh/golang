package main

import "fmt"

func main() {

	x := bar()

	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
}

func bar() func() int {
	return func() int {
		return 1992
	}
}
