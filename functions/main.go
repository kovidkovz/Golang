package main

import (
	"fmt"
)

// the func main in go acts as an entry point to our code, this is the reason why we dont need to call this main func to execute this function.

func main() {
	fmt.Println("welcome to functions on golang")

	// call the greet function
	greet()

	a := addition(1, 2)

	b := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(a)
	fmt.Println("Sum using proadder is:", b)

	fact, msg, name := fact(4, "Kovid")
	fmt.Println("factorial:", fact, msg, name)

	p := isPrime(8497498)
	fmt.Println(p)
}

// this function does not return anything
func greet() {
	fmt.Println("Hello everyone I am a function")
}

// this is a function that will take arguments and at last return a value
// In a function that returns a value, we need to define the datatype of the value that is bieng returned by the function
// while passing the arguments we need to define the datatype of the arguments
func addition(num1 int, num2 int) int {
	return num1 + num2
}

// This is an example of a function that takes multiple arguments, the '...' in the arguments show that the funtion takes multiple arguments.
func proAdder(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

// We can return more that one value as well. We need to just define the datatypes of the values that we want to return
func fact(num int, name string) (int, string, string) {
	f := 1
	for i := 1; i <= num; i++ { //traditional for loop
		f *= i
	}

	return f, "wohoooo!!", name
}

// function to calculate whether is prime or not
func isPrime(num int) string {
	if num <= 1 {
		return "undefined"
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return "not prime"
		}
	}
	return "prime"
}
