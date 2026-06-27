package main

import (
	"fmt"
	"time"
)

func Cook(dish string) {
	fmt.Println("Cooking:", dish)
	time.Sleep(2 * time.Second)
	fmt.Println("Finished cooking:", dish)
}

func main() {
	fmt.Println("Coook...")
	go Cook("Dinner")
	go Cook("Lunch")
	go Cook("Breakfast")
	time.Sleep(3 * time.Second)
	fmt.Println("Finish this...")
}
