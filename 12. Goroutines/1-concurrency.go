package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0)) // shows total cpu cores of current machine
	// Concurency = A chef cooking multiple dishes, rapidly switching back and forth in preparing each dish
	// Parallelism = Three separate chefs cooking one dish each
}
