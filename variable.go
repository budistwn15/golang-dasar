package main

import "fmt"

func main() {
	var name string
	name = "Spy x family"
	fmt.Println(name)

	name = "Spy"
	fmt.Println(name)

	var fullName = "Spy x Family"
	fmt.Println(fullName)

	var age int8 = 17
	fmt.Println("Umur saya adalah ", age)

	gender := "female"
	fmt.Println(gender)

	var (
		firstName = "Spy"
		lastName  = "x Family"
	)

	fmt.Println(firstName, lastName)
}
