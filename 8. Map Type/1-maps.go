package main

import "fmt"

func sellFood(food map[string]int, foodType string, quant int) {
	currentFood, exists := food[foodType]

	if exists {
		if currentFood >= quant {
			fmt.Println("You ordered:", quant, foodType)
			food[foodType] -= quant
			fmt.Println("There is/are currently", food[foodType], foodType, "left.")
		} else {
			fmt.Println("Not enough", foodType, "available.")
		}
	} else {
		fmt.Println(currentFood, "is not available. Please try again soon!")
	}
}

func main() {
	menu := map[string]int{
		"Breakfast": 100,
		"Lunch":     150,
		"Dinner":    200,
	}

	for key, value := range menu {
		fmt.Println("Key:", key, "| Value:", value)
	}

	menu["Snack"] = 50

	fmt.Println(menu)

	// Delete keys from map
	delete(menu, "Snack")

	fmt.Println(menu)

	// Check if key exists
	exists := menu["Snack"]
	fmt.Println(exists)

	// var stock map[int]string - zero value is nil
	stock := make(map[string]int)

	stock["Breakfast"] = 11
	stock["Dinner"] = 12

	fmt.Println(stock)

	sellFood(stock, "Breakfast", 10)
	sellFood(stock, "Dinner", 13)
}
