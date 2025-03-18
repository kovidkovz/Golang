package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://example.com:8000/path?param=value"

func main() {
	fmt.Println("welcome to handling the urls")

	result, _ := url.Parse(myurl)
	fmt.Println("result's scheme is:", result.Scheme)
	fmt.Println("Host is:", result.Host)
	fmt.Println("Path is:", result.Path)
	fmt.Println("Raw query is:", result.RawQuery)
	fmt.Println("Port is:", result.Port())

	qparams := result.Query()
	fmt.Printf("The type of qparams is %T\n", qparams)

	fmt.Println("value is", qparams["param"])
}
