package main

import (
	"encoding/json"
	"fmt"
)

type Agent struct {
	Name   string
	Age    int
	Dept   string
	Kills  int
	Ops    int
	Ltk    bool
	Armory []string
}

func main() {

	agent1 := Agent{
		Name:   "Jitendar",
		Age:    30,
		Dept:   "Feild Agent",
		Kills:  56,
		Ops:    5,
		Ltk:    true,
		Armory: []string{"M416", "DP-28", "KAR-98", "P226"},
	}

	agent2 := Agent{
		Name:   "Sandip",
		Age:    33,
		Dept:   "Operational Agent",
		Kills:  0,
		Ops:    15,
		Ltk:    true,
		Armory: []string{"P226"},
	}

	// var SecretAgent []Agent

	SecretAgent := []Agent{agent1, agent2}

	for _, v := range SecretAgent {
		fmt.Println(v.Name, v.Age, v.Dept, v.Kills, v.Ops, v.Ltk)
		for _, sv := range v.Armory {
			fmt.Println(sv)
		}
	}

	// fmt.Println(p1)
	// people := []Person{p1, p2}
	// fmt.Println(people)

	bs, err := json.Marshal(SecretAgent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
