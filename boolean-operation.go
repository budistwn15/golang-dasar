package main

import "fmt"

func main() {
	var (
		price      = 5000
		qty        = 10
		isDiscount = price >= 3000 && qty >= 4
	)

	fmt.Println(isDiscount)

}
