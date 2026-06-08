package main 

import "fmt"

func main() {
	fruit := "Banana"

	fmt.Println(fruit)
	fmt.Println("Memory location:", &fruit)
	fmt.Printf("Pointer address: %p\n", &fruit)

	fruitCopy := fruit
	fmt.Println("Memory location:", &fruitCopy)
}
