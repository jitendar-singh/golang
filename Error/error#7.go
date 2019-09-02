package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string
	Age     int
	Sayings []string
}

func main() {

	p1 := Person{
		Name: "jitendar Singh",
		Age:  27,
		Sayings: []string{
			"Shaken not Stirred",
			"Never Say never",
		},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON did not Marshal - ", err)
	}
	fmt.Println(string(bs))

}
