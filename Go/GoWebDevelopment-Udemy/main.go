package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	fmt.Fprintf(w, "This is the About Page\nSum = %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32
	x, y = 100, 0
	f, err := DivideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is equal to %f", x, y, f))
}

func DivideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by 0!")
		return 0, err
	}
	result := x / y
	return result, nil
}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe(":8686", nil)
}
