package main

import "fmt"

var myName string

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Good morning"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to: ", i)

	whatWasSaid1, whatWasSaid2 := saySomething()
	fmt.Println(whatWasSaid1, whatWasSaid2)
}

func saySomething() (string, string) {
	return "something", "else"
}
