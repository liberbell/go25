package main

import "log"

func main() {
	// var myString string
	// var myInt int

	// myString = "hi"
	// myInt = 11

	// mySecondString := "another string"

	// log.Println(myString, myInt, mySecondString)

	myMap := make(map[string]int)

	// myMap["dog"] = "Eric"
	// myMap["other-dog"] = "Bob"
	// myMap["dog"] = "Alex"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	myMap["First"] = 1
	myMap["Second"] = 2

	log.Panicln(myMap["First"], myMap["Second"])
}
