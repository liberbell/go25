package main

import (
	"fmt"
	"net/http"

	_ "github.com/liber/myapp/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlrs.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
