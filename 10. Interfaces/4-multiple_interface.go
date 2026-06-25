package main

import "fmt"

type BoxerInfo interface {
	GetInfo() string
}

type BoxerFight interface {
	Fight(f string) string
}

type Boxer struct {
	Name   string
	Fights int
}

func (b *Boxer) GetInfo() string {
	return fmt.Sprintf("Boxer's Name: %s | Fights: %d", b.Name, b.Fights)
}

func (b *Boxer) Fight(f string) string {
	b.Fights += 1
	return fmt.Sprintf("Boxer: %s | Vs | Boxer: %s", b.Name, f)
}

func main() {
	fv := Boxer{"Manny", 73}

	var info BoxerInfo = &fv
	var fight BoxerFight = &fv

	fmt.Println(info.GetInfo())
	fmt.Println(fight.Fight("Luis"))
	fmt.Println(info.GetInfo())

}
