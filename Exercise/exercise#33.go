package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

var bs []byte

func main() {

	u1 := user{
		First: "Jitendar",
		Age:   30,
	}

	u2 := user{
		First: "Monalisa",
		Age:   28,
	}

	people := []user{u1, u2}

	bs, error := json.Marshal(people)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(bs))

}
