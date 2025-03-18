package main

import "fmt"

// slices are powerful as compared to arrays in golang
func main() {
	// declare a slice
	fruitList := []string{"Apple", "Banana", "Corn"}

	fruitList = append(fruitList, "mango", "juice")

	fmt.Println("this is a fruitlist:", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println((fruitList))

	fruitList = append(fruitList[1:3]) // here the element at index 3 is not included
	fmt.Println((fruitList))

	highscores := make([]int, 4)

	highscores[0] = 1
	highscores[1] = 2
	highscores[2] = 3
	highscores[3] = 4

	fmt.Println("highscores:", highscores)
	
	// when we use append all the items inside the slice go through reallocation of memory to accomodate the appended items in the slice.
	highscores = append(highscores, 12, 13, 14)

	fmt.Println("appended highscores:", highscores)

	// remove items from a slice 
	languages := []string{"Python", "Java", "Ruby", "Swift", "Javascript"}
	var index = 2
	languages = append(languages[:index], languages[index+1:]...) //here the ... are used to define the definition of a function that accepts more arguments
	fmt.Println("sliced slice:", languages)
}


// basic difference in declaring a slice and an array
// array []string || [size]string{a,b,c} || [size]string
// slice []string{} || []string{a,b,c}