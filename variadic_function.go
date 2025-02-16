package main

import "fmt"

func sumAll(numbers ...int) int  {
	total := 0;
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
 result := sumAll(10, 20, 30, 40, 50)

 fmt.Println(result)
 fmt.Println(sumAll(10, 10, 10))
 fmt.Println(sumAll(100, 50, 10))

 numbers := []int{10, 10, 10, 10}
 fmt.Println(sumAll(numbers...))
}