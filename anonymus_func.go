package main

import "fmt"

func main() {

	func() {
		fmt.Println(`This is my first anonymus function`)
	}()
	func(x int, name string) {
		fmt.Println(name, ` lets use the value of x in anonymus function 2 : `, x)
	}(5, `Jitendar`)
}
