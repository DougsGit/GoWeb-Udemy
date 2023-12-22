package main

import (
	"fmt"
	"net/http"
	"github.com/DougsGit/GoWeb-Udemy/Go/GoWebDevelopment-Udemy/templates"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
