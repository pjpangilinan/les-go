package main

import "fmt"

func main() {
	var err error
	err = fmt.Errorf("This is an error.")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error encountered.")
	}

	panic(err)
}
