package main

import (
	"fmt"
	"io"
	"net/http"
)

const webpage string = "https://www.codecademy.com/"

func main() {
	fmt.Println("Welcome to handling web requests.")
	response, err := http.Get(webpage)

	if err != nil {
		panic(err)
	}

	fmt.Printf("This is the type of the response %T\n", response)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("This is the data inside the body:", string(body))
}