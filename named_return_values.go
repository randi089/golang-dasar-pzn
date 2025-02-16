package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string)  {
	firstName = "Randi"
	middleName = "Febriadi"
	lastName = "Hebat"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName, middleName, lastName)
}