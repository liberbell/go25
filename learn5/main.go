package main

import "log"

type myStruct struct {
	Firstname string
}

func (m *printFirstName) string {
	
}

func main() {
	var myVar myStruct
	myVar.Firstname = "Bob"

	myVar2 := myStruct{
		Firstname: "Eric",
	}

	log.Println("myVar is set to ", myVar.Firstname)
	log.Println("myVar2 is set to ", myVar2.Firstname)
}
