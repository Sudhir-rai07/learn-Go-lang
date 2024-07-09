package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time")

	// Time is valuable
	var currentTime time.Time = time.Now() // Gives current time
	fmt.Println("Current time : ", currentTime)
	formattedTime := currentTime.Format("02/01/2006 monday 15:04") // Formates the date with silly standard format setups
	fmt.Println("Formatted time : ", formattedTime)

}
