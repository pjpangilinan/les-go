package main

import "fmt"

func ten_second_timer() (func(step int) int, int) {
	countDown := 10

	var stepDown = func(step int) int {
		countDown -= step
		return countDown
	}

	return stepDown, countDown
}

func main() {
	stepDown, countDown := ten_second_timer()
	fmt.Println("The countdown starts at", countDown)

	fmt.Println(stepDown(1))
	fmt.Println(stepDown(1))
	fmt.Println(stepDown(1))
	fmt.Println(stepDown(3))

}
