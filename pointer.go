package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Padang Pariaman", "Sumatera Barat", "Indonesia"}
	// address2 := address1 // copy value
	// address2.City = "Padang"

	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah menjadi Padang

	var address1 Address = Address{"Padang Pariaman", "Sumatera Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Padang"

	fmt.Println(address1) // ikut berubah menjadi Padang
	fmt.Println(address2) // berubah menjadi Padang
}
