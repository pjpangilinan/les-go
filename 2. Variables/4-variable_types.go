package main

import "fmt"

func main() {
	productName := "Cheerios"
	price := 30.00
	ready := true
	count := 31

	fmt.Printf("Type of productName is: %T\n", productName)
	fmt.Printf("Type of price  is: %T\n", price)
	fmt.Printf("Type of ready is: %T\n", ready)
	fmt.Printf("Type of count is: %T\n", count)
}
