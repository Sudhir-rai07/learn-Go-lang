package main

import "fmt"

func main() {
	fmt.Println("Welcomr to Struct")
	// No inheritance in Golang; No super or parent
	user := User{"Naina", "naina@gmail", true, 24}

	fmt.Println((user.Name))
	fmt.Printf("User is %+v\n", user)
	fmt.Printf("User name is %v and email is %v", user.Name, user.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
