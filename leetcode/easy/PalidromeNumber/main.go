package main
import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	firstString := strconv.Itoa(x)
	reverseString := ""
	for _, char := range firstString {
		reverseString = string(char) + reverseString
	}
	fmt.Println(firstString, reverseString)	
	return strings.Compare(firstString, reverseString) == 0
}

func main() {
	fmt.Println(isPalidrome(121))
}
