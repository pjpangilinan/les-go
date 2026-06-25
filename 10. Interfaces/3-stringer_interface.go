package main

import "fmt"

type Order struct {
	Customer string
	ItemName string
	Quant    int
	Cost     int
}

func (o Order) String() string {
	return fmt.Sprintf("Order: %s has ordered %d %s for a total of $%d.", o.Customer, o.Quant, o.ItemName, o.Cost*o.Quant)
}

func main() {
	order := Order{"Hodge", "Air Fryer", 12, 20}
	fmt.Println(order)
}
