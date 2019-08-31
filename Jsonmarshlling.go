package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string
	Age      int
	Fav_food []string
}

func main() {

	p1 := Person{
		Name:     "Jitendar Singh",
		Age:      30,
		Fav_food: []string{"Chicken", "Mutton", "Palak Paneer"},
	}

	p2 := Person{
		Name:     "Sandip Singh",
		Age:      33,
		Fav_food: []string{"Chicken", "Mutton", "Palak Paneer"},
	}

	// fmt.Println(p1)
	people := []Person{p1, p2}
	// fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
