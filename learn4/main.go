package main

import (
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
}

func main() {
	user := User{
		FirstName: "Bob",
		LastName:  "Mary",
	}
}
