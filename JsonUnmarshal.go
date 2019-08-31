package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"Name"`
	Age     int      `json:"Age"`
	FavFood []string `json:"Fav_food"`
}

func main() {
	s := `[{"Name":"Jitendar Singh","Age":30,"Fav_food":["Chicken","Mutton","Palak Paneer"]},
	{"Name":"Sandip Singh","Age":33,"Fav_food":["Chicken","Mutton","Palak Paneer"]}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []Person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data ", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.Name, v.Age)
		for _, n := range v.FavFood {
			fmt.Println(n)
		}
	}

}
