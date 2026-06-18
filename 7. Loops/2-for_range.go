package main

import "fmt"

func main() {
	ratings := []float64{1.23, 5.00, 3.44, 4.13, 2.50}

	var sum float64

	for i, rating := range ratings {
		fmt.Printf("Rating #%d: %.2f\n", i+1, rating)
		sum += rating
	}

	avg := sum / float64(len(ratings))

	fmt.Printf("Average Rating: %.2f\n", avg)
}
