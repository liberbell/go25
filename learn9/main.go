package main

import "github.com/helpers"

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber

}

func main() {
	// log.Println("Hello")

	// var myVar helpers.SomeType
	// myVar.TypeName = "some name"
	// log.Println(myVar.TypeName)
	intChan := make(chan int, 0)
}
