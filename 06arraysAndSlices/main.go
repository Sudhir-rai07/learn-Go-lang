package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays")

	// Array initialization
	// var marksArr = [5]int{4, 5, 6, 7, 911} // If you don't pass length of array in brackets then it will behave like "slices"
	// fmt.Println(marksArr)
	// fmt.Println("Length of Array marks is: ", len(marksArr))

	fmt.Println("Slice")
	var marksSlice = []int{4, 5, 6, 7, 911, 112, 412}
	fmt.Println(marksSlice)
	fmt.Println("Length of Slice marks is: ", len(marksSlice))

	// fmt.Println(marksSlice[1:])                 // [5 6 7 911 112 412]
	// fmt.Println(marksSlice[:len(marksSlice)-1]) // [4 5 6 7 911 112]
	// fmt.Println(marksSlice[2:4]) // [6 7]

	marksSlice = append(marksSlice, 90, 92, 32)
	fmt.Println("New marks slice : ", marksSlice)
	fmt.Println("Length of Slice marks is: ", len(marksSlice))

	var newSlice []string
	fmt.Printf("len of s %v cap of s %v\n", len(newSlice), cap(newSlice))
	if newSlice == nil {
		fmt.Println("Nil newSlice")
	}

}
