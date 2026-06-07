package main

import "fmt"

func main() {

	greet()

	var sum = addition(5, 6)

	fmt.Println(sum)

	var diff = subtraction(10, 2)

	fmt.Println(diff)

	// Named function literal
	var multiply = func(num1, num2 int) int {
		return num1 * num2
	}

	fmt.Println(multiply(2,3))
}

func greet() {
	fmt.Println("Hello there, welcome to here")
}

func addition(num1 int, num2 int) int {
	return (num1 + num2)
}

// Named return value
func subtraction(num1 int, num2 int) (difference int) {
	difference = num1 - num2
	return
}
