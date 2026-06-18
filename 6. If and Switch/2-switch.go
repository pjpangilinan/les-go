package main

import "fmt"

func main() {
	today := "Thursday"

	switch today {
	case "Monday", "Tuesday", "Wednesday":
		fmt.Println("The weekends are far away.")
	case "Thursday", "Friday":
		fmt.Println("The weekends are so close.")
	case "Saturday", "Sunday":
		fmt.Println("Its the weekend!")
	default:
		fmt.Println("I don't know what day it is today.")
	}
}
