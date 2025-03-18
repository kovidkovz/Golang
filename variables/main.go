package main

import (
	"fmt"
)

// define constant, a public variable starts with an uppercase
const Logintoken string = "ifnwiqfni9wvfuiwnd"

func main() {
	var name string = "John Cena"
	println(name)
	var age int8 = 56
	fmt.Printf("The age of %v is %v years\n", name, age)
	fmt.Printf("%T \n", age)

	var isLoggedIn bool = true
	fmt.Printf("The isLoggedIn variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 //untyped integer max val 255 only
	fmt.Printf("The smallVal variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.7789247235476419412758126475612479085642 //untyped integer max val 255 only
	fmt.Println(smallFloat)
	fmt.Printf("The smallFloat variable is of type: %T \n", smallFloat)

	var number int = 100000000000000
	fmt.Println(number)
	fmt.Printf("The number variable is of type: %T \n", number)

	// implicit type, 
	var website = "www.example.com"
	fmt.Printf("The variable website is of type: %T \n", website)

	// only inside the function cannot be used for global or public variable declaration.
	dog := "Bark"
	fmt.Printf("The variable dog is of type: %T \n", dog)
}
