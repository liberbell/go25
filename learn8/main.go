package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOFTeeth int
}

func main() {
	dog := Dog{
		Name:  "Eric",
		Breed: "German Shepherds",
	}

	PrintInfo(dog)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal syas", a.Says(), " and has", a.NumberOfLegs(), "legs.")
}

func (d Dog) Says() string {
	return "woof"

}
