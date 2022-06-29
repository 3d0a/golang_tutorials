package main
import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	splitedString :=  strings.Split(s, " ")
	return len(splitedString[len(splitedString)-1])
}

func main() {

}
