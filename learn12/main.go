package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page.")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValue(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the about page. and 2 + 2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot devide by zero")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f diveided %f is %f", 100.0, 0.0, f))
}

func addValue(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil
}

func divideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
