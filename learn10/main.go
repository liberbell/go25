package main

import (
	"encoding/json"
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
			"first_name: "Bob",
			"last_name": "Mary",
			"hair_color": "brown",
			"has_dog": false,
		},
		{
			"first_name": "Elton",
			"last_name": "John",
			"hair_color": "gold",
			"has_dog": true,
		}
	]`

	var unmarshaled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshaled)
	if err != nil {
		log.Println("Error unmarshalled json", err)
	}
	log.Printf("unmarshaled: %v", unmarshaled)
}
