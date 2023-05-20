package main

import "log"

func main() {
	// var myString string
	// var myInt int

	// myString = "hi"
	// myInt = 11

	// mySecondString := "another string"

	// log.Println(myString, myInt, mySecondString)

	myMap := make(map[string]string)

	myMap["dog"] = "Eric"
	myMap["other-dog"] = "Bob"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])
}
