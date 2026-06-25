package main

import "fmt"

type Cleaner interface {
	Clean() string
}

type JuniorCleaner struct {
	Name string
}

func (j JuniorCleaner) Clean() string {
	return fmt.Sprintf("The building has been cleaned by Junior Cleaner, %s.", j.Name)
}

type SeniorCleaner struct {
	Name string
}

func (s SeniorCleaner) Clean() string {
	return fmt.Sprintf("The building has been cleaned by Senior Cleaner, %s.", s.Name)
}

func CheckClean(c Cleaner) {
	fmt.Println(c.Clean())
	fmt.Println("Building has been cleaned.")
}

func main() {
	var jose, klay Cleaner
	jose = SeniorCleaner{"Jose"}
	klay = JuniorCleaner{"Klay"}

	CheckClean(jose)
	CheckClean(klay)

	// Promote Klay
	klay = SeniorCleaner{"Klay"}
	CheckClean(klay)
}
