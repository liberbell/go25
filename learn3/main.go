package main

import "log"

var s = "seven"

func main() {
	var s2 = "six"
	var s string

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
