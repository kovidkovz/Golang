package main

import "fmt"

func main() {
	fmt.Println("This is how maps look like in golang")
	// here we have to define the datatype of the key as well as the value
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	defer fmt.Println("languages are:", languages)

	fmt.Println("JS is short for:", languages["JS"])
	fmt.Println("PY is short for:", languages["PY"])
	fmt.Println("RB is short for:", languages["RB"])

	keys := make([]string, 0, len(languages))
	for k := range languages {
		keys = append(keys, k)
	}
	fmt.Println("keys:", keys)

	delete(languages, "RB")

	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value)
	}

	for key, _ := range languages{
		fmt.Println(languages[key])
	}
	
	languages["SDK"] = "software development kernel"
}
