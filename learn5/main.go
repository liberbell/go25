package main

import "log"

type myStruct struct {
	Firstname string
}

func (m *myStruct) printFirstName() string {
	return m.Firstname
}

func main() {
	var myVar myStruct
	myVar.Firstname = "Bob"

	myVar2 := myStruct{
		Firstname: "Eric",
	}

	log.Println("myVar is set to ", myVar.Firstname)
	log.Println("myVar2 is set to ", myVar2.Firstname)

	log.Println("myVar is set to ", myVar.printFirstName())
	log.Println("myVar2 is set to ", myVar2.printFirstName())
}
