package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 100000
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai64)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var (
		name    = "Spy x family"
		e       = name[0]
		stringE = string(e)
	)

	fmt.Println(stringE)
}
