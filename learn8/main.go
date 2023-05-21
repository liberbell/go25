package main

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
}
