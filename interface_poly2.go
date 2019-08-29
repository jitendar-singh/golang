package main

import (
	"fmt"
	// "time"
)

type Skill struct {
	ltk bool
}

type Agent struct {
	Skill
	first  string
	last   string
	drinks []string
}

type Person struct {
	Agent
	Skill
}

func main() {

	p1 := Person{

		Agent{
			first: `Jitendar`,
			last:  `Singh`,
			drinks: []string{
				`Vodka`,
				`Martini`,
				`Beer`,
			},
		}, Skill{
			ltk: true,
		},
	}

	for _, v := range p1.drinks {
		fmt.Println(v)
	}
}


g9hdvq5w@Sunny@987