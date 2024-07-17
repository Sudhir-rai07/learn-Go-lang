package main

import "fmt"

func main() {

	fmt.Println("Welcome to defer in golang")
	// Without defer
	fmt.Println("World")
	fmt.Println("Hello")
	fmt.Println("-------")

	// With defer
	defer fmt.Println("World")
	fmt.Println("Hello")

	fmt.Println("Without defer")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("With defer")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
