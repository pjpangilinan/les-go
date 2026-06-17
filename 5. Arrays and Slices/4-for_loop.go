package main

import "fmt"

func main() {
	countDown := []int{10, 9, 8, 7, 6, 5}

	for index, num := range countDown {
		fmt.Printf("Count: %d : Down: %d\n", index, num)
	}
}
