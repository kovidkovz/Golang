package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1

	switch dicenumber {
	case 1:
		fmt.Println("open up")
	
	case 2:
		fmt.Println("move 2 points")
	
	case 3:
		fmt.Println("move 3 points")

	case 4:
		fmt.Println("move 4 points")

	case 5:
		fmt.Println("move 5 points")
	
	case 6:
		fmt.Println("move 6 points")
	}
}
