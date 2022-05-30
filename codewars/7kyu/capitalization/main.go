package main
import "fmt"

func Capitalize(st string) []string {
  resultArray  := make([]string, 0)
  firstString  := make([]rune, 0)
  secondString := make([]rune, 0)
  for index, char := range st {
    if index % 2 == 0 {
	  firstString  = append(firstString, char - rune(32))
	  secondString = append(secondString, char)
	} else {
		secondString  = append(secondString, char - rune(32))
		firstString = append(firstString, char)
	}  
  }
  resultArray = append(resultArray, string(firstString), string(secondString))
  return resultArray
}

func main() {
  fmt.Println(Capitalize("abcdefghijklmnopqrstuvwxyz"))
}