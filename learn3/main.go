package main

import (
	"log"
	"time"
)

var s = "seven"
var firstName string
var lastName string
var age int
var birthDate time.Time

func main() {
	var s2 = "six"

	s := "eight"

	log.Println("s  is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")

	log.Println("after call func s is", s)
	log.Println("s2 is", s2)
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saysomething func is", s)
	return s3, "world"
}
