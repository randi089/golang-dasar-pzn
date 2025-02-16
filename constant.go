package main

import "fmt"

func main() {
	// const firstName string = "Randi"
	// const lastName = "Febriadi"
// ---------------------------------------------------
	const (
		firstName string = "Randi"
		lastName = "Febriadi"
	) 

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Fauzan"
	// lastName = "Fatwa"
}