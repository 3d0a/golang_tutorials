package main
import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string{
	finalArray := make([]string, 0)
	stringArray := strings.Split(str, " ")
	for _, word := range stringArray {
		reversedWord := ""
		for _, char := range word {
			reversedWord = string(char) + reversedWord
		}
		finalArray = append(finalArray, reversedWord)
	}
	return strings.Join(finalArray, " ")
}

func main() {
	fmt.Println(ReverseWords("hello fucking world!"))
}

