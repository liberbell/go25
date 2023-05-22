package main

import (
	"log"

	"github.com/liber/myniceprogram/helpers"
)

const numPool = 1000

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
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
