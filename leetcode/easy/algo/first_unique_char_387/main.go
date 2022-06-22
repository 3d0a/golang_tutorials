package main

import (
	"fmt"
	"strings"
)

func firstUniqChar(s string) int {
	for i, char := range s {
		if strings.Count(s, string(char)) == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
