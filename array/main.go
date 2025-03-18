package main

import "fmt"

func main() {
	fmt.Println("This is how an array looks like in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[3] = "Orange"
	fruitList[2] = "Banana"

	fmt.Println("here is the fruitlist:", fruitList)

	var vegList = [4]string{"Potato", "Onion", "Brinjal"}
	fmt.Println("veglist:", vegList)
	// thats pretty much all from arrays,there is no sorting as well in arrays
	fmt.Printf("type of veglist is %T\n", vegList)
}
