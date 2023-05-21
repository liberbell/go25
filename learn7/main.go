package main

import (
	"log"
)

func main() {
	// var isTrue bool
	// isTrue = false

	// if isTrue == true {
	// 	log.Println("isTrue is ", isTrue)
	// } else {
	// 	log.Println("isTrue is ", isTrue)
	// }

	// cat := "cat2"

	// if cat == "cat" {
	// 	log.Println("Cat is cat")
	// } else {
	// 	log.Println("Cat is not cat")
	// }

	// myNum := 100
	// isTrue := false

	// if myNum > 99 && !isTrue {
	// 	log.Println("Mynum is greater than 99 and istrue is set to true")
	// } else if myNum < 100 && isTrue {
	// 	log.Println("1")
	// } else if myNum == 100 || isTrue {
	// 	log.Println("2")
	// } else if myNum > 1000 && isTrue {
	// 	log.Println("3")
	// }
	myVar := "fish"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	}
}
