package main

import (
	"fmt"
	"strings"
)

func UniqChars(s string) bool {
	s = strings.ToLower(s)
	seen := make(map[rune]bool)

	for _, ch := range s {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

func main() {
	fmt.Println(UniqChars("abcd"))
	fmt.Println(UniqChars("abCdefAaf"))
	fmt.Println(UniqChars("aabcd"))
}
