package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	var hpwd, pwd, temppwd []byte

	fmt.Printf("\n %T \t %T\n ", hpwd, pwd)

	pwd = []byte("htc@03569Ya")

	hpwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	// Assume you fetched the pwd from the database & stored in temppwd
	temppwd = []byte("htc@03569Ya")
	fmt.Println(hpwd, temppwd)

	err = bcrypt.CompareHashAndPassword(hpwd, temppwd)
	if err != nil {
		fmt.Println("You cannot Login")
		return
	}
	fmt.Println("You are logged in")

}
