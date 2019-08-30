package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) Speak() {
	fmt.Println(" My name is ", p.first, "", p.last, " and my age is ", p.age)
}
func main() {

	p1 := Person{
		first: "Jitendar",
		last:  "Singh",
		age:   29,
	}

	p1.Speak()

}
