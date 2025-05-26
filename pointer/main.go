package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	names := []string{"Kovid", "Aman", "Riya"}
	students := []*Student{}

	for _, name := range names {
		s := Student{Name: name}
		students = append(students, &s)
	}

	for _, student := range students {
		fmt.Println(student.Name)
	}
}
