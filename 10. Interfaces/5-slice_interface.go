package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Customer struct {
	Name string
}

func (c Customer) Greet() string {
	return fmt.Sprintf("Hello, how are you? My name is %s.", c.Name)
}

type Staff struct {
	Name string
	Role string
}

func (s Staff) Greet() string {
	return fmt.Sprintf("Welcome to our lovely shop! My name is %s and I'll be your %s for today.", s.Name, s.Role)
}

func main() {
	greeters := []Greeter{
		Customer{"Garuda"},
		Staff{"Jam", "Barista"},
		Customer{"Baruuk"},
	}

	for _, g := range greeters {
		fmt.Println(g.Greet())
	}

	greeters = append(greeters, Staff{"Pepe", "Cashier"})

	for _, g := range greeters {
		fmt.Println(g.Greet())
	}
}
