package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Randi"
	names[1] = "Febriadi"
	names[2] = "Hebat"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// var values = [3]int {
	// 	90,
	// 	80,
	// 	95,
	// }

	var values = [...]int {
		90,
		80,
		95,
		100,
		110,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}