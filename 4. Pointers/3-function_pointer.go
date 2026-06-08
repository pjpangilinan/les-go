package main

import "fmt"

func applyDiscount(originalPrice *float64, disc float64) {
	*originalPrice = *originalPrice - (*originalPrice * disc)
	fmt.Println(&originalPrice)
}

func main() {
	fruitPrice := 20.00
	discount := 0.10
	fmt.Println(&fruitPrice)
	fmt.Printf("Original fruit price: %.2f\n", fruitPrice)

	applyDiscount(&fruitPrice, discount)
	fmt.Printf("Discounted fruit price: %.2f\n", fruitPrice)
}
