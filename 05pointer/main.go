package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var num int = 20           // Value stored at any location of memory
	var addOfNum *int = &num   // Address of that memory
	var number int = *addOfNum // Value at that memory address

	fmt.Println("Address of num : ", addOfNum)
	fmt.Println("Number at address of num : ", number)
	fmt.Println("Number at address of num : ", number*2)

}
