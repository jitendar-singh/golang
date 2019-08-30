package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := sumAll(s1...)
	fmt.Println(`Sum of all : `, sum)

	sum2 := even(sumAll, s1...)
	fmt.Println(`Sum of all even : `, sum2)

	sum3 := odd(sumAll, s1...)
	fmt.Println(`Sum of odd `, sum3)

}

func sumAll(x ...int) int {

	sum := 0

	for _, v := range x {
		sum += v
	}
	return sum

}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var zi []int

	for _, v := range vi {
		if v%2 != 0 {
			zi = append(zi, v)
		}
	}
	return f(zi...)
}
