package main

import "fmt"

func main() {
	var (
		a = 10
		b = 10
		d = 5
		e = 2
		c = a + b - d * e
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	
	i += 5 // i = i + 5
	fmt.Println(i)

	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}