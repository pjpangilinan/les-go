package main

import "fmt"

func main() {
	price := 50.50
	qty := 3
	total := price * float64(qty)

	fmt.Printf("Total price: %.2f\n", total)
}
