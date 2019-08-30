package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(xi...)
	fmt.Println(sum)

	xii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum2 := bar(xii)
	fmt.Println(sum2)
}

func foo(x ...int) int {
	var sum int

	for _, v := range x {
		sum = sum + v
	}
	return sum

}

func bar(xi []int) int {
	var sum int

	for _, v := range xi {
		sum = sum + v
	}
	return sum

}
