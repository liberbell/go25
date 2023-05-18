package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Bob",
		LastName:    "Mary",
		PhoneNumber: "111-222-3333",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate:", user.BirthDate)
}
