package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("Welcome to webrequests in golang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response = %T\n", response)

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(data)
	fmt.Printf("Type of content = %T\n", content)
	fmt.Println(content)
}
