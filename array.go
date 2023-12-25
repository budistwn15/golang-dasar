package main

import "fmt"

func main() {
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Mango"

	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])

	var prices = [3]int{
		150,
		120,
		100,
	}

	fmt.Println(prices)
	fmt.Println(len(fruits), len(prices))
}
