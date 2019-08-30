package main

import "fmt"

func main() {

	a := 10
	fmt.Println("a before ", a)
	fmt.Println("a before ", &a)
	foo(&a)
	fmt.Println("a after ", a)
	fmt.Println("a after ", &a)

}

func foo(x *int) {
	fmt.Println("x before ", x)
	fmt.Println("x before ", *x)
	*x++
	fmt.Println("x after ", x)
	fmt.Println("x after ", *x)
}
