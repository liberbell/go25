package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Bob",
		"last_name": "Mary",
		"hair_color": "brown",
		"has_dog": false
	},
	{
		"first_name": "Elton",
		"last_name": "John",
		"hair_color": "gold",
		"has_dog": true
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalled json", err)
	}
	log.Printf("unmarshaled: %v", unmarshalled)

	var mySlice []Person
	var m1 Person
	m1.FirstName = "Alex"
	m1.LastName = "Hepp"
	m1.HairColor = "black"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m1.FirstName = "Eric"
	m1.LastName = "Crapton"
	m1.HairColor = "gray"
	m1.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("error marshalling", err)
	}
	fmt.Println(string(newJson))
}
