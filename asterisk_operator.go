package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Padang Pariaman", "Sumatera Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Padang"

	fmt.Println(address1) // ikut berubah menjadi Padang
	fmt.Println(address2) // berubah menjadi Padang

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // address 1 berubah
	fmt.Println(address2)
}
