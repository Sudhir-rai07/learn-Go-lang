package main

import "fmt"

type Student struct {
	Name  string
	Title string
}

func (fn Student) fullName() string {
	return (fn.Name + " " + fn.Title)
}

// Same Same but diffrent
func fullName01(fn Student) string {
	return (fn.Name + " " + fn.Title)
}

// Method with pointer
func (fn *Student) pfullName() string {
	fn.Name = "Hitesh"
	return fn.Name
}

func main() {
	fmt.Println("Welcome to methods")

	fName := Student{"Ruhi", "Sharma"}
	fmt.Println(fName.fullName())
	fmt.Println(fullName01(fName))
	fmt.Println(fName)
	fmt.Println(fName.pfullName())
	fmt.Println(fName)
}
