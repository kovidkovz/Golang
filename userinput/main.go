package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// declare a reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma,ok || err ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
}
