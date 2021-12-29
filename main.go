package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("This is the Home Page"))
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addFun(2, 3)
	fmt.Fprintf(w, fmt.Sprintf("This is the About Page abd sum of 2 nd 3 is %d", sum))
}

func addFun(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in dividing"))
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("Divided value of %f %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by 0")
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
