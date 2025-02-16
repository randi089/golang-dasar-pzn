package main

import "fmt"

func main() {
	name := "Randi"

	switch name {
	case "Randi":
		fmt.Println("Hello Randi")
	case "Fatwa":
		fmt.Println("Hello Fatwa")
	case "Fauzan":
		fmt.Println("Hello Fauzan")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

	name = "Febriadi"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}