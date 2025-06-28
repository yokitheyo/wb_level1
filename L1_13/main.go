package main

import "fmt"

func main() {
	a, b := 55, 163
	fmt.Println("before:", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("after:", a, b)
}
