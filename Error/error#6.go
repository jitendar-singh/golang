package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("Hello.txt")
	if err != nil {
		log.Println("Error happened", err)
	}
	defer f2.Close()

	fmt.Println("Check the log.txt file in the directory")
}
