package main

import "fmt"

type Customer struct {
	Name string
	Age int
}

func (customer Customer) greetCustomer() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	first := Customer{"Brandon", 15}

	first.greetCustomer()
}


