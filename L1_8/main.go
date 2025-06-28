package main

import "fmt"

func setBit(num int64, i uint, bit int) int64 {
	if bit == 1 {
		return num | (1 << i)
	} else {
		return num &^ (1 << i)
	}
}

func main() {
	var num int64 = 5
	fmt.Printf("original number: %d (bin %04b)\n", num, num)

	num = setBit(num, 1, 1)
	fmt.Printf("after setting bit 1 to 1: %d (bin %04b)\n", num, num)

	num = setBit(num, 3, 1)
	fmt.Printf("after setting bit 3 to 1: %d (bin %04b)\n", num, num)
}
