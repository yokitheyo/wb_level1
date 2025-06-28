package main

import "fmt"

func binarySearch(arr []int, target int) int {
	lhs, rhs := 0, len(arr)-1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			lhs = mid + 1
		} else {
			rhs = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println("index 7:", binarySearch(arr, 7))
	fmt.Println("index 10:", binarySearch(arr, 10))
}
