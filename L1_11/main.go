package main

import "fmt"

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 3, 3, 4}
	res := make([]int, 0)

	set := make(map[int]struct{})

	for _, v := range A {
		set[v] = struct{}{}
	}

	seen := make(map[int]struct{})
	for _, v := range B {
		if _, ok := set[v]; ok {
			if _, already := seen[v]; !already {
				res = append(res, v)
				seen[v] = struct{}{}
			}
		}
	}

	fmt.Println(res)
}
