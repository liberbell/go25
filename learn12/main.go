package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page.")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := AddValue(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is", sum))
}

func AddValue(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)
}
