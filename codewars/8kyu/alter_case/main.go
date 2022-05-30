package main
import (
	"fmt"
	"strings"
)

func ToAlternatingCase(str string) string {
	resultSlice := make([]string, 0)
	for _, char := range str {
		if char >= 65 && char <= 90 {
			char += 32
			resultSlice = append(resultSlice, string(char))
		}  else if char >= 97 && char <= 122 {
			char -= 32
			resultSlice = append(resultSlice, string(char))
		}  else {
			resultSlice = append(resultSlice, string(char))
		}
	}
	return strings.Join(resultSlice, "")
}

func main() {
	fmt.Println(ToAlternatingCase("String.prototype.toAlternatingCase123"))
}

