package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var randi Customer
	fmt.Println(randi)

	randi.Name = "Randi Febriadi"
	randi.Address = "Indonesia"
	randi.Age = 26

	fmt.Println(randi)
	fmt.Println(randi.Name)
	fmt.Println(randi.Address)
	fmt.Println(randi.Age)

	fatwa := Customer{
		Name:    "Khairul Fatwa",
		Address: "Medan",
		Age:     25,
	}
	fmt.Println(fatwa)

	fauzan := Customer{"Muhammad Fauzan", "Payakumbuh", 26}
	fmt.Println(fauzan)

	fauzan.sayHello("Randi")
	randi.sayHello("Fatwa")
	fatwa.sayHello("Randi")
}
