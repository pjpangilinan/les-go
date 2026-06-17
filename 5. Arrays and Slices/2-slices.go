package main

import "fmt"

func main() {
	drinkSizes := [4]string{"Small", "Venti", "Large", "Grande"}

	fmt.Println(drinkSizes)

	sliceSize := drinkSizes[0:2]
	// sliceSize = drinkSizes[:] - all elements
	// sliceSize = drinkSizes[2:] - [Large Grande]
	// sliceSize = drinkSizes[:2] - [Small Venti]

	fmt.Println(sliceSize)
	fmt.Println(sliceSize[1])
	fmt.Println(sliceSize[0])

	sliceSize[0] = "Extra Small" // Changing slice also changes array value

	fmt.Println(drinkSizes)
	fmt.Println(sliceSize)

	// slices without arrays
	ratings := []int{1, 2, 3, 4, 5}
	cTypes := make([]string, 3)

	fmt.Println(ratings, cTypes)

}
