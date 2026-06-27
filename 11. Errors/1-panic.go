package main

import "fmt"

func ServeFood(food string, amount int, customers int) {
	fmt.Printf("Serve %d %s to %d customers.\n", amount, food, customers)
	amountPerCustomer := amount / customers
	fmt.Println(amountPerCustomer)
}

func BetterServeFood(food string, amount int, customers int) {
	/*
		if amount < 1 || customers < 1 {
			fmt.Println("Amount or Customers cannot be zero.")
			return
		} */

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	fmt.Printf("Serve %d %s to %d customers.\n", amount, food, customers)
	amountPerCustomer := amount / customers
	fmt.Println(amountPerCustomer)
}

func main() {
	BetterServeFood("Banana", 12, 6)
	BetterServeFood("Banana", 12, 0)
	// ServeFood("Banana", 12, 0) // panic: runtime error: integer divide by zero
}
