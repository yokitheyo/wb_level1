package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	var lhs, rhs []int
	for _, v := range arr[1:] {
		if v <= pivot {
			lhs = append(lhs, v)
		} else {
			rhs = append(rhs, v)
		}
	}
	return append(append(quickSort(lhs), pivot), quickSort(rhs)...)
}

func main() {
	arr := []int{10, 5, 2, 3, 7, 6, 1, 8}
	sorted := quickSort(arr)
	fmt.Println("sorted:", sorted)
}
