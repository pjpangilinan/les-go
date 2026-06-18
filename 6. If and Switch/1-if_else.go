package main

import "fmt"

func main() {
	number := 10
	lessNumber := 10

	if number > lessNumber {
		fmt.Println("Good number:", number, "| Less number :", lessNumber)
	} else if number < lessNumber {
		fmt.Println("Good number:", lessNumber, "| Less number :", number)
	} else if number == lessNumber {
		fmt.Println("Same number.")
	} else {
		fmt.Println("No number.")
	}

	ranks := [5]string{"Bronze", "Silver", "Gold", "Platinum", "Diamond"}

	score := 78
	level := 3

	// Combined condition (AND)
	if score >= 50 && level >= 2 {
		fmt.Println("Eligible for upgrade check")
		// Nested if
		if score >= 90 {
			fmt.Println("Exceptional performance")
			if level >= 4 {
				fmt.Println("Top-tier player detected")
			} else {
				fmt.Println("High score but level not maxed")
			}
		} else if score >= 70 {
			fmt.Println("Good performance")
			if ranks[level] == "Gold" || ranks[level] == "Platinum" {
				fmt.Println("Already in advanced rank:", ranks[level])
			} else {
				fmt.Println("Eligible for rank promotion")
			}
		} else {
			fmt.Println("Average performance")
		}
	} else {
		fmt.Println("Not eligible for upgrade check")
	}

	var multiply = func(num1 int, num2 int) int {
		return (num1 * num2)
	}

	if product := multiply(2, 3); product > 10 {
		fmt.Printf("Product is %d - Good\n", product)
	} else {
		fmt.Printf("Product is %d - Bad\n", product)
	}
}
