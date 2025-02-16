package main

import "fmt"

func getHello(name string) string  {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Randi")
	fmt.Println(result)

	fmt.Println(getHello("Fatwa"))
	fmt.Println(getHello("Fauzan"))
}