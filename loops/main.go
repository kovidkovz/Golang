package main

import "fmt"

func main() {
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println("Weekdays:", weekdays)

	// for i := range weekdays{
	// 	fmt.Println(weekdays[i])
	// }

	// for index, day := range weekdays{
	// 	fmt.Printf("At index %v value is %v\n", index, day)
	// }

	// for _, day := range weekdays{
	// 	fmt.Println(day)
	// }

	threshold := 1

	
	for threshold < 10 {

	if threshold ==4 {
		break
	}
	
	fmt.Println(threshold)
	threshold++
	}

}