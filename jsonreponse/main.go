package main

import (
	"encoding/json"
	"fmt"
)

type Orders struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Price   int    `json:"price,omitempty"`
	Country string `json:"country,omitempty"`
}

func main() {
	fmt.Println("Json response")

	resp := []Orders{
		{1, "item1", 230, "India"},
		{2, "item2", 450, "Australia"},
		{3, "item3", 560, "America"}, 
	}

	fmt.Println("resp:", resp)

	byteresp, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", byteresp)

	var custresponse []Orders

	err = json.Unmarshal(byteresp, &custresponse)
	if err != nil {
		panic(err)
	}

	fmt.Println("bytes:", custresponse)

	fmt.Println(custresponse[0].Country)

	for count, value := range custresponse {
		fmt.Println("value", value)
		fmt.Println("count", count)
	}
}
