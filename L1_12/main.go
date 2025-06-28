package main

import "fmt"

func main() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, animal := range animals {
		set[animal] = struct{}{}
	}

	for animal := range set {
		fmt.Println(animal)
	}
}
