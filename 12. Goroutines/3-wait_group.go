package main

import (
	"fmt"
	"sync"
	"time"
)

func Cook(dish string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Cooking:", dish)
	time.Sleep(1 * time.Second)
	fmt.Println("Finished cooking:", dish)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Coook...")
	wg.Add(3)
	go Cook("Dinner", &wg)
	go Cook("Lunch", &wg)
	go Cook("Breakfast", &wg)

	wg.Wait()

	// time.Sleep(3 * time.Second)

	fmt.Println("Finish this...")
}
