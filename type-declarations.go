package main

import "fmt"

func main() {
	type married bool
	var isMarried married = true

	fmt.Println("Apakah kamu sudah menikah? ", isMarried)
}
