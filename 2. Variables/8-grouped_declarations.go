package main

import "fmt"

func main() {
	var (
		name string = "Me"
		age int = 45
		netWorth float64 = 215461.23
	)
	fmt.Printf("My name is %s, I am %d years old. I am worth %.2f pesos.\n", name, age, netWorth)
}
