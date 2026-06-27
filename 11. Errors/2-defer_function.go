package main

import "fmt"

func Greet(message string, receipient string) {
	fmt.Println("This is a message from: ", receipient)
	fmt.Println(message)
}

func main() {
	defer Greet("I dont want this.", "Me")
	fmt.Println("I like this.")

	defer fmt.Println("I dislike this.")

	defer func(asd int) {
		fmt.Println("Hello there!")
		fmt.Println(asd)
	}(12)

	fmt.Println("I like this. 2")

}
