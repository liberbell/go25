package main

import "log"

func main() {
	var mySlice []int
	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 1)

	log.Println(mySlice)
}
