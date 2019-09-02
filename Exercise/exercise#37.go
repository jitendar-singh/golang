package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("hello")
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p1 := person{
		first: "James",
	}
	saySomething(&p1)

	// saySomething(p1)
	// does not work

	p1.speak()

}
