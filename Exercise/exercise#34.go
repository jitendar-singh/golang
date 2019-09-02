package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	raw_json := []byte(`[{"First":"Jitendar","Age":30},{"First":"Monalisa","Age":28}]`)

	people := []user{}

	error := json.Unmarshal(raw_json, &people)
	if error != nil {
		fmt.Println(error)
	}
	for k, v := range people {
		fmt.Println("User", k+1)
		fmt.Println("Name:", v.First)
		fmt.Println("Age:", v.Age)
		fmt.Println("-------------------")
	}
}
