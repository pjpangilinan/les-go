package main

import "fmt"

func main() {
	slice := []string{"pop", "del", "slice", "app"}

	fmt.Println(slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	// Add elements to slice via append

	slice = append(slice, "depend")

	fmt.Println(slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	slice = append(slice, "asd")

	fmt.Println(slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	// Addresses are the same for the slices above

	sos := make([]string, 10, 100)
	sos[0] = "asda"
	fmt.Println(sos)

	// Delete via append, first argument decides which element to delete via index
	// In this case slice[:1], deletes the second element which is del
	indexRemove := 1
	slice = append(slice[:indexRemove], slice[indexRemove+1:]...)
	// slice[indexRemove+1:]... is equal to slice[2], slice[3], slice[4], slice[5]
	fmt.Println(slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))
	// Deletion can also be done by import "slices" and using Delete.
	// Same arguments like this - <slice>.Delete(<slice>[:indexRemove], <slice>[indexRemove+1:]...)

}
