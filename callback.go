/* Passing a func as an argument */
package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(ii...))
	sum1 := even(sum, ii...)
	fmt.Println(sum1)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total = total + v
	}
	return total
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
