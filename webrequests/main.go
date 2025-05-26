package main

import (
	"fmt"
	"io"
	"net/http"
)


const url = "https://www.testprepreview.com/modules/reading1.htm"

func main() {
	fmt.Println("welcome to webrequests")

	response, err :=http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("The type of reponse is %T\n", response)

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println("Content:", content)
} 
