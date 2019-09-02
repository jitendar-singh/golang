package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// This will throw an error .
	// As there is no file named notes in the directory
	f, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("This is a file open Error- ", err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))

}
