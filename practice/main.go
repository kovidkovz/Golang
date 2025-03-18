package main

import "fmt"

// variables
// private(unexported variable), starts with a lowercase and can be used within the same package only not in other packages
var name = "John Cena"

// public(exported variable), starts with an uppercase and can be used in other packages as well and are not restricted to only the package where they are created 
var Name = "John Cena"


func main() {
	fmt.Println(name)
	fmt.Println(Name)

	// only to be declared within a function
	age := 4

	fmt.Println(age)
}