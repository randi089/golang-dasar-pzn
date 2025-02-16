package main

import "fmt"

func main() {
	type NoKTP string

	var (
		ktpRandi NoKTP = "1111111111"

		contoh string = "222222222"
		contohKtp NoKTP = NoKTP(contoh)
	)

	fmt.Println(ktpRandi)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
}