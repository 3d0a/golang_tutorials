package main
import (
  "fmt"
  "strconv"
)

func SumOfIntegersInString(str string) int {
  resultArr := make([]int, 0)
  digit := ""
  for i, char := range []rune(str) {
	if char >= '0' && char <= '9' {
	  digit = digit + string(char)
	  if i == len(str) - 1 {
	    semiResultInt, _ := strconv.Atoi(digit)
		resultArr = append(resultArr, semiResultInt)
	  }
	} else if len(digit) != 0 {
	  semiResultInt, _ := strconv.Atoi(digit)
	  resultArr = append(resultArr, semiResultInt)
	  digit = "" 
	}
  }
  max := 0
  for _, digit := range resultArr {
    max += digit
  }
  return max
}

func main() {
  fmt.Println(SumOfIntegersInString("12.4"))
}