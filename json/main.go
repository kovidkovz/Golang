package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	p := Person{Name :"Alice", Age: 12,Address:"India"}

	

	defer ConvertTostruct()

	response, err:= json.Marshal(p)

	if err != nil {
		panic(err)
	}

	fmt.Println("This is the json reponse:", string(response))
	fmt.Println(response)

	
}

// this is how a json Marshal method is used
// We use the Marshall method to convert the data or struct into a json format
// Similarly, we use the unmarshal method to convert a json format back to the type struct.

// I have shown how to convert a json to struct

// I will create a new function that will take json data as an input and try to perform unmarshalling
// What a superb way to convert a json format into a struct. isn't it.
// jai om namah shivaya thankyou god for everything...... you are the best of all
// Always be with us o lord the supreme power, you are our king and we are your slaves. 
// Always give us your blessings and love we are your devotees
// let the devil away from us
// we pray to you o our king

func ConvertTostruct() {
	jsonData := `{"name":"Alice","age":12,"address":"India"}`

	var p Person
	err := json.Unmarshal([]byte(jsonData), &p) //This returns only one value

	if err!= nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(p.Name, p.Age, p.Address)

}


func (p *Person) Walk() {
	fmt.Println("a man is walking")
}