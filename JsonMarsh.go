package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
	Ltk   bool
}

func main() {

	p1 := Person{
		First: "Jitendar",
		Last:  "Bond",
		Age:   27,
		Ltk:   true,
	}

	p2 := Person{
		First: "Monalisa",
		Last:  "Moneypenny",
		Age:   27,
		Ltk:   false,
	}

	agent := []Person{
		p1,
		p2,
	}

	bs, err := json.Marshal(agent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
