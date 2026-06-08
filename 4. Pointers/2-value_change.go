package main

import "fmt"

func main() {
	fruitPrice := 500

	fmt.Println("Memory address:", &fruitPrice)

	var pointerFruitPrice *int = &fruitPrice
	*pointerFruitPrice = 200

	fmt.Println("Updated fruitPrice:", fruitPrice)
}
