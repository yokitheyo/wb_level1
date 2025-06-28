package main

import "fmt"

func reverseBytes(s []byte, lhs, rhs int) {
	for lhs < rhs {
		s[lhs], s[rhs] = s[rhs], s[lhs]
		lhs++
		rhs--
	}
}

func reverseWords(s string) string {
	bytes := []byte(s)
	n := len(bytes)

	reverseBytes(bytes, 0, n-1)

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || bytes[i] == ' ' {
			reverseBytes(bytes, start, i-1)
			start = i + 1
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
}
