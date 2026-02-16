package main

import (
	"fmt"
	"time"
)

func main() {
	baton := make(chan int, 5)
	go func() {
		time.Sleep(3* time.Second)
		fmt.Println("waiting for data")
		for i := range baton{
			fmt.Println("data:", i)
		}
	}()

	// time.Sleep(2* time.Second)
	// fmt.Println("sending data")
	for i:=1; i<=9; i++{
		fmt.Println("sending data")
		baton <- i
		fmt.Println("Sent")
	}
	close(baton)
	// fmt.Println("Sent")
	// time.Sleep(2* time.Second)
	
}
