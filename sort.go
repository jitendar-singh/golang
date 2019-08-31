package main

import (
	"fmt"
	"sort"
)

func main() {

	var x []int
	var str []string

	x = []int{10, 8, 99, 45, 7, 54, 32, 3, 2, 87}
	str = []string{"Sunny", "Runny", "Chi-Chi", "Kiara"}
	fmt.Println(x)
	sort.Ints(x)
	fmt.Println(x)

	fmt.Println(str)
	sort.Strings(str)
	fmt.Println(str)
}
