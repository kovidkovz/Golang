package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float32
	Job    string
	Email  string
}

func main() {
	kovid := User{"Kovid", 23, 90000.00, "SDE", "kovidbhatt27@gmail.com"}
	kovid.newname()
	fmt.Printf("%+v\n",kovid)
	fmt.Println(kovid)
	kovid.getstatus()
	
}

// these are the methods of a struct
func (u User) getstatus(){
	fmt.Println("The user's email is:", u.Email)
}

// Pointer's are used to point to the value at the actual address, and not create a copy of that variable
func (u *User) newname(){
	u.Name = "Kovid Bhatt"
	fmt.Println("user's new name is:", u.Name)
}