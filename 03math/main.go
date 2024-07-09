package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Welcome to math class")
	var maxNum int = math.MaxInt64 // max value of integer
	fmt.Println("Max value of integer in go is : ", maxNum)
	var minNum int = math.MinInt64 // max value of integer
	fmt.Println("Min value of integer in go is : ", minNum)

	num := 5.0
	squrtONum := math.Sqrt(num)
	pow2Onum := math.Pow(num, 2)
	fmt.Println("Square root of number is : ", squrtONum)
	fmt.Println("Power 2 of number is : ", pow2Onum)

	// many more as per usages

}
