package main

import "fmt"

// Empty Interface
type Any interface{}

type Enclosure interface {
	Clean() string
	Feed() string
}

type LionEnclosure struct {
	Name string
	Food string
}

type TigerEnclosure struct {
	Name string
	Food string
}

func (e LionEnclosure) Clean() string {
	return fmt.Sprintf("%s the lion's enclosure has been cleaned.", e.Name)
}

func (e LionEnclosure) Feed() string {
	return fmt.Sprintf("%s the lion has been fed.", e.Name)
}

func (e TigerEnclosure) Clean() string {
	return fmt.Sprintf("%s the tiger's enclosure has been cleaned.", e.Name)
}

func (e TigerEnclosure) Feed() string {
	return fmt.Sprintf("%s the tiger has been fed.", e.Name)
}

func main() {
	var lion, tiger Enclosure
	lion = LionEnclosure{"Mushi", ""}
	fmt.Println(lion.Clean())
	fmt.Println(lion.Feed())

	tiger = TigerEnclosure{"Libre", ""}
	fmt.Println(tiger.Clean())

	// Empty interface can be assigned any value type
	var any Any = "Hello"
	fmt.Println(any)
	any = []string{"H", "e", "l"}
	fmt.Println(any)
	any = 15
	fmt.Println(any)

	anyType := []interface{}{
		"popo",
		12.5,
		false,
		[1]string{"Hello"},
		[]int{1, 2, 3, 4},
		6,
	}

	for _, t := range anyType {
		fmt.Println(t)
	}

}
