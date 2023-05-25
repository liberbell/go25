package main

import (
	"net/http"
)

func main() {

	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {})
}
