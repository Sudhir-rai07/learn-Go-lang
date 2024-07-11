package main

import "fmt"

func main() {
	fmt.Println("Welcome to go maps")

	marks := make(map[int]int)
	marks[0] = 98
	marks[1] = 99
	marks[2] = 90
	marks[3] = 88

	fmt.Println("List of marks : ", marks)
	fmt.Println("0 : ", marks[0])

	delete(marks, 0)
	fmt.Println("List of marks ", marks)

	for i := 1; i <= len(marks)+1; i++ {
		fmt.Println(marks[i] + 5)
	}
}
