package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to conditionals pt-2 : Switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can go 1 foreward")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can add 3 spots")
	case 4:
		fmt.Println("It's 4, Go four steps foreward")
	case 5:
		fmt.Println("Go +5")
		fallthrough
	case 6:
		fmt.Println("Hurrah!! It's 6. play again.")
	}

}
