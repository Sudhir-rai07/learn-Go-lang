package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in go")

	// Fo has only for loop and it's variations

	// for i := 1; i < 11; i++ {
	// 	// Continue
	// 	if i == 7 {
	// 		continue
	// 	}

	// 	// Break
	// 	if i == 9 {
	// 		fmt.Println("Loop breaked")
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	//01
	list := []string{"Hello", "Lisa", "baby", "How", "are", "You"}
	for l := 0; l < len(list); l++ {
		fmt.Println(list[l])
	}

	// 02 The range form of the for loop iterates over a slice or map.
	for index, value := range list {
		fmt.Printf("Index %v and value %v\n", index, value)
	}

	//03 The range form of the for loop iterates over a slice or map.: Without index value
	for _, value := range list {
		fmt.Printf("Value %v\n", value)
	}

	// While loop
	var i int = 0
	for i < 10 {
		fmt.Printf("Number is %v\n", i)
		i++
	}
}
