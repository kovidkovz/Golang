package main

import "fmt"

func main() {
	count := 4

	if count != 4 {
		fmt.Println("count is 4")
	} else if count <4 {
		fmt.Println("count is less than 4")
	} else {
		fmt.Println("count is greater than 4")
	}

	if count%2 ==0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
