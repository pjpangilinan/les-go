package main

import "fmt"

func main() {
	for i := 10; i > 0; i-- {
		fmt.Println("Starting in :", i)
	}

	num := 10
	for num > 0 {
		fmt.Println("Number:", num)
		num--
	}

	for {
		var name string
		fmt.Println("Type your name or 'exit' to quit.")
		fmt.Scanln(&name)

		if name == "exit" {
			break
		} else {
			fmt.Println("Hello there", name)
		}
	}
}
