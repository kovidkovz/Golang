package main

import (
	"bufio"
	"fmt"
	"os"
)

// calculator

func main() {
	fmt.Println("Welcome to my calculator")

	values := [4]float32{}
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.Read()

}

func multiplication(values ...float64) float64 {
	product := 1.0

	for _, val := range values {
		product *= val
	}

	return product
}
