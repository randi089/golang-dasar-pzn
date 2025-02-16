package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Randi"
	// person["address"] = "Padang Pariaman"

	person := map[string]string{
		"name": "Randi",
		"address": "Padang Pariaman",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Randi"
	book["ups"] = "Salah"

	fmt.Println(book)
	
	delete(book, "ups")
	
	fmt.Println(book)
}