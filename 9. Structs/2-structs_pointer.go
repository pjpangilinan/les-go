package main

import "fmt"

func orderDone(order *FoodOrder) {
	order.Status = true
	fmt.Printf("Thanks for ordering a %s, %s. That will cost %d dollars.\n", order.FoodType, order.CustomerName, order.Cost)
}

type FoodOrder struct {
	FoodType     string
	Cost         int
	CustomerName string
	Status       bool
}

func main() {
	third := FoodOrder{"Snack", 20, "Kyle", false}
	orderDone(&third)

	fmt.Println(third)
}
