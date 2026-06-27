package main

import (
	"fmt"
	"time"
)

func PrepareDish(ord string, orderC chan string) {
	fmt.Println("Preparing your dish:", ord)
	time.Sleep(2 * time.Second)
	fmt.Printf("The chef has finished preparing your %s.\n", ord)
	orderC <- ord
}

func main() {
	orderC := make(chan string)

	fmt.Println("Welcome to our restaurant!")

	var ord string
	fmt.Println("What is your order?")
	fmt.Scan(&ord)

	go PrepareDish(ord, orderC)

	ord = <-orderC
	fmt.Printf("Your %s is ready. Please come get it.\n", ord)

	fmt.Println("The restaurant is now closed.")
}
