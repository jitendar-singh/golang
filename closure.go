package main

import "fmt"

func main() {

	a := incrementer()
	b := incrementer()
	fmt.Printf("%T", a)
	fmt.Println(b)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementer() func() int {
	var x int
	return func() int {
		x = x + 1
		return x
	}

}
