package main

import "fmt"

type filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello", filteredName)
// }

func sayHelloWithFilter(name string, filter filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Randi", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
