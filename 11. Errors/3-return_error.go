package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("notes.txt")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File doesn't exist.")
		} else {
			fmt.Println("Error encountered", err)
		}
		return
	}

	fmt.Println("File opened:", file.Name())
}
