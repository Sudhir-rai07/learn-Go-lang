package main

import "fmt"

func main() {
	fmt.Println("Welcomr to Struct")
	// No inheritance in Golang; No super or parent
	user := User{"Naina", "naina@gmail", true, 24}

	fmt.Println((user.Name))
	fmt.Printf("User is %+v\n", user)
	fmt.Printf("User name is %v and email is %v\n\n", user.Name, user.Email)

	// A slice can have a type of struct
	result := []Result{
		{90, true},
		{88, true},
		{67, true},
		{45, true},
		{28, false},
	}
	fmt.Println(result)
}

type Result struct {
	marks int
	pass  bool
}
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
