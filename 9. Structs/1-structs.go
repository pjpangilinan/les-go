package main

import "fmt"

func printReceipt(order FoodOrder) {
	fmt.Printf("Thanks for ordering a %s, %s. That will cost %d dollars.\n", order.FoodType, order.CustomerName, order.Cost)
}

type FoodOrder struct {
	FoodType     string
	Cost         int
	CustomerName string
	Status       bool
}

func main() {
	var first FoodOrder

	first.FoodType = "Breakfast"
	first.Cost = 100
	first.CustomerName = "Brandon"

	fmt.Println(first)

	// Struct literals
	second := FoodOrder{
		FoodType:     "Dinner",
		Cost:         200,
		CustomerName: "Sheila",
	}

	// inline struct literal - second := FoodOrder{"Dinner", 200, "Sheila"}
	printReceipt(second)
}
