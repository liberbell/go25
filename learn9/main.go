package main

import (
	"log"

	"github.com/liber/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)
}
