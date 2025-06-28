package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	arr := []int{2, 4, 6, 8, 10}

	for _, num := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(num * num)
		}()
	}
	wg.Wait()
}
