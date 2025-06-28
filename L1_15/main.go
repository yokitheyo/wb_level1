package main

import "strings"

var justString string

// используется всего 100 байт из 1024 байт
func someFunc() {
	v := createHugeString(1 << 10)
	// копируем только 100 байт
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
