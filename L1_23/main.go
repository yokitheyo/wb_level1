package main

import "fmt"

func RemoveI[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		panic("index out of range")
	}

	copy(slice[i:], slice[i+1:])

	var zero T
	slice[len(slice)-1] = zero
	return slice[:len(slice)-1]
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	strings := []string{"a", "b", "c", "d"}

	ints = RemoveI(ints, 2)
	strings = RemoveI(strings, 1)

	fmt.Println("ints:", ints, len(ints), cap(ints))
	fmt.Println("strings:", strings, len(strings), cap(strings))
}
