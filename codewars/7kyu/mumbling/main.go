package main 
import (
	"fmt"
	"strings"
)

func Accum(s string) string {
  arrOfStrings := make([]string, 0)
  for i, char  := range []rune(s) {
	if (char >='a' && char <= 'z') || ((char >='A' && char <= 'Z')){
	  stringChar  := string(char)
	  arrOfStrings = append(arrOfStrings, strings.ToUpper(stringChar) + strings.Repeat(strings.ToLower(stringChar), i))
	} else {
		return "0"
	}
  }
  return strings.Join(arrOfStrings, "-")
}

func main() {
  fmt.Println(Accum("ZpglnRxqenU"))
}