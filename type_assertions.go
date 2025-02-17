package main

import "fmt"

func random() any {
	return true
}

func main() {
	var (
		result any = random()
		// resultString string = result.(string)
	)
	// fmt.Println(resultString)

	// Error atau Panic
	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("String", value)
	default:
		fmt.Println("Unknown", value)
	}
}
