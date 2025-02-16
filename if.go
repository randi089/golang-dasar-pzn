package main

import "fmt"

func main() {
	name := "Randi"

	if name == "Randi" {
		fmt.Println("Hello Randi")
	} else if name == "Fatwa" {
		fmt.Println("Hello Fatwa")
	} else if name == "Fauzan" {
		fmt.Println("Hello Fauzan")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}