package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	fmt.Println((*p).first)
	fmt.Println((*p).last)
	fmt.Println((*p).age)
}
func main() {
	p1 := person{
		first: "Jitendar",
		last:  "Singh",
		age:   27,
	}
	changeMe(&p1)

}
