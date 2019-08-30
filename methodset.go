package main

import "fmt"

type SecretAgent struct {
	first string
	last  string
	age   int
}

type Person struct {
	SecretAgent
	ltk bool
}

type Agent interface {
	database()
}

func main() {

	p1 := Person{
		SecretAgent: SecretAgent{
			first: "Jitendar",
			last:  "Bond",
			age:   27,
		},
		ltk: true,
	}

	p2 := Person{
		SecretAgent: SecretAgent{
			first: "James",
			last:  "Singh",
			age:   30,
		},
		ltk: false,
	}
	p1.database()
	p2.database()
}
func (c *Person) database() {
	fmt.Println(c.first)
	fmt.Println(c.last)
	fmt.Println(c.age)
	fmt.Println(c.ltk)
}
