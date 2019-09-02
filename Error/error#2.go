package main

import "fmt"

func main() {

	var ans1, ans2, ans3 string

	fmt.Print("Name: ")
	_, err := fmt.Scanf("%s", &ans1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scanf("%s", &ans2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	// This will throw am Error
	_, err = fmt.Scanf("%d", &ans3)
	if err != nil {
		panic(err)
	}

	fmt.Println(ans1, ans2, ans3)

}
