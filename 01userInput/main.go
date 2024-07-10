package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// User Input
	fmt.Println("Enter a number")

	// Text reader
	reader := bufio.NewReader(os.Stdin)

	// text Input
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Converts the string to number
	num, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Text = ", text)
	fmt.Printf("Type of text before conversion = %T\n", text)
	fmt.Println("After conversion")
	fmt.Println("Text = ", num)
	fmt.Printf("Type of text after conversion = %T\n", num)
	num++
	fmt.Printf("Incrementing the number  = %v", num)

}
