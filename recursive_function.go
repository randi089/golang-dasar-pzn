package main

import "fmt"

func factorial(angka int) int {
	if angka > 1 {
		return angka * factorial(angka-1)
	} else {
		return 1
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	resultFunction := factorial(5)
	resultLoop := factorialLoop(5)

	fmt.Println(resultFunction)
	fmt.Println(resultLoop)
}
