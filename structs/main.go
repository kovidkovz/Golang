package main

import "fmt"

type Animal struct {
	name    string
	sound   string
	colour  string
	extinct bool
	age     int
}

func main() {
	fmt.Println("This is how a struct looks like in golang")

	dog := Animal{"Bruno", "Bark", "Black", false, 8}

	fmt.Println("This is a dog:", dog)
	fmt.Printf("The dog struct %+v\n", dog)

	fmt.Println(dog.name)
	fmt.Printf("dog: %v\n", dog)

	// we can create a slice of type strut we just created

	animal_list := []Animal{
		{"Deagon", "chirp", "grey", true, 30},
		{"Dinosaur", "bradnva", "green", true, 100},
	}

	fmt.Println("this is an animal list:", animal_list)

	fmt.Println(dog.age)
	
}