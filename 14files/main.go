package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to files")

	text := "Hello, I am writing this text in a file. Hope it'll be saved."

	file, err := os.Create("./myFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err) // --> Both are right
	length, err := io.WriteString(file, text)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	fmt.Println("Length is : ", length)

	fmt.Println(file)
	defer file.Close()

	readFile("./myFile.txt")

}

func readFile(filename string) {

	databyte, err := os.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println(string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
