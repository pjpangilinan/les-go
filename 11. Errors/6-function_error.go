package main

import "fmt"

type OutofStock struct {
	Item string
}

func (o OutofStock) Error() string {
	return fmt.Sprintf("%s is out of stock", o.Item)
}

func PresentItem(item string) (string, error) {
	quant := items[item]
	if quant == 0 {
		return "", OutofStock{Item: item}
	}
	items[item]--
	return fmt.Sprintf("Here is your item: %s", item), nil
}

func IsError(r string, e error) string {
	if e != nil {
		return fmt.Sprintf("Error encountered: %s.", e)
	} else {
		return r
	}
}

var items = map[string]int{
	"Snack":  12,
	"Drinks": 5,
	"Meal":   0,
}

func main() {
	receipt, err := PresentItem("Meal")
	fmt.Println(IsError(receipt, err))

	receipt, err = PresentItem("Drinks")
	fmt.Println(IsError(receipt, err))

	receipt, err = PresentItem("Drinks")
	fmt.Println(IsError(receipt, err))

	receipt, err = PresentItem("Drinks")
	fmt.Println(IsError(receipt, err))

	receipt, err = PresentItem("Drinks")
	fmt.Println(IsError(receipt, err))

	receipt, err = PresentItem("Drinks")
	fmt.Println(IsError(receipt, err))

	receipt, err = PresentItem("Drinks")
	fmt.Println(IsError(receipt, err))

}
