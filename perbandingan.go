package main

import "fmt"

func main() {
	var (
		name1 = "Randi"
		name2 = "Randi"

		// result bool = name1 == name2 // true
		result bool = name1 != name2 // false
	)

	fmt.Println(result)
}