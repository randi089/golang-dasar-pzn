package main

import "fmt"

type HasName interface {
	getName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Randi"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
