package main

import "fmt"

func main() {
	var someFruits = [3]string{"Banana", "Orange", "Apple"}

	fmt.Println(someFruits)
	fmt.Println(len(someFruits))
	fmt.Println(len("Hello World!"))

	someFruits[2] = "Grapes"

	fmt.Println(someFruits[0])
	fmt.Println(someFruits[1])
	fmt.Println(someFruits[2])
}
