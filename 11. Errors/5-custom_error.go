package main

import "fmt"

type YourError string

func (y YourError) Error() string {
	return string(y)
}

func main() {
	var err error
	err = YourError("This is an error.")

	if err != nil {
		fmt.Println("Error encountered:", err)
	} else {
		fmt.Println("Continue on...")
	}
}
