package main

import (
	"errors"
	"fmt"
)

func CheckR(radi int) error {
	if radi > 1000 {
		return errors.New("Radiation exceeds safety parameters. Please evacuate immediately.")
	} else if radi > 500 {
		return fmt.Errorf("Radiation critical at %d rads. Employ safety measures immediately.", radi)
	} else {
		return nil
	}
}

func main() {
	radiations := [4]int{10, 800, 400, 1200}

	for _, rad := range radiations {
		err := CheckR(rad)
		if err != nil {
			fmt.Println("Error encountered:", err)
		} else {
			fmt.Printf("Nuclear reactor running at %d rads within safety parameters. Continue with monitoring.\n", rad)
		}
	}
}
