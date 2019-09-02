package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("Sunny.txt")
	if err != nil {
		fmt.Println("File Open Error", err)
		// log.Println("Err happened", err)
		log.Fatalln(err)
		// panic(err)
		return
	}
	defer f.Close()
}
