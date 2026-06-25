package main

import "fmt"

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Address   string

	FullName func(customer Customer)
}

func getFullName(customer Customer) {
	fmt.Println(customer.FirstName, customer.LastName)
}

func main() {
	var first = Customer{"Brandon", "Lee", 26, "Monty Street 2516", getFullName}

	fmt.Println(first)
	first.FullName(first)
}
