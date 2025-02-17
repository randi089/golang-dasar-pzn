package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	randi := Man{"Randi"}
	randi.Married()

	fmt.Println(randi.Name)
}
