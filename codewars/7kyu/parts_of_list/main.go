package main
import (
  "fmt"
  "strings"
)

func PartList(arr []string)  {
  resultArr := make([]string, 0)
  for i := 1; i < len(arr); i++ {
	semiResult := make([]string, 0)
    semiResult = append(semiResult, strings.Join(arr[0:i], " "))
	semiResult = append(semiResult, strings.Join(arr[i:len(arr)], " "))
	resultArr  = append(resultArr,   strings.Join(semiResult, ", ")) 
  }
  resultString := strings.Join(resultArr, ")(")
  resultString = "(" + resultString + ")"
  fmt.Println(resultString)
}

func main() {
  PartList([]string{"az", "toto", "picaro", "zone", "kiwi"})
}