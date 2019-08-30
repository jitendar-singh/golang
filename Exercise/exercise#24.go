package main

import "fmt"

func main() {
	f1()
	defer f2()
	f3()

}

func f1() {
	fmt.Println("This is f1")
}
func f2() {
	fmt.Println("This is f2")
}
func f3() {
	fmt.Println("This is f3")
}
