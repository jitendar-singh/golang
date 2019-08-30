package main

import "fmt"

func main() {

	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int
	b = &a
	fmt.Printf("The value of b %v\n", b)
	fmt.Printf("The value b is pointing at  %v\n", *b)

	*b = 43
	fmt.Printf("The value of a %v value of b %v\n", &a, *b)
}
