package main

import "fmt"

func main() {
	fmt.Println("Welcome to conditionals(if-else)")

	var count int = 8

	if count > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Greate part didn't run")
	}

	if count < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Less part didn't run")
	}

	//
	if num := 10; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Not less than 10")
	}
}
