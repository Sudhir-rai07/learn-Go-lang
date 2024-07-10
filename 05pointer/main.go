package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var num int = 20
	var addOfNum *int = &num
	var number int = *addOfNum

	fmt.Println("Address of num : ", addOfNum)
	fmt.Println("Number at address of num : ", number)
	fmt.Println("Number at address of num : ", number*2)

}
