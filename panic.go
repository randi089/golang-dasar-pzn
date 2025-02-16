package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Randi Febriadi Hebat")
}
