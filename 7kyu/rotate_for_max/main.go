package main
import (
  "fmt"
  "strings"
  "strconv"
  "sort"
)

func MaxRot(n int64) int64 {
  baseString := strings.Split(strconv.FormatInt(n, 10), "")
  resultArr := make([]string, 0)
  resultArr = append(resultArr, strconv.FormatInt(n, 10))
  for i := 0; i<len(baseString); i++ {
    for j := i; j < len(baseString) - 1; j++ {
      baseString[j], baseString[j+1] = baseString[j+1], baseString[j]
	  }
    resultArr = append(resultArr, strings.Join(baseString, ""))
  }
  sort.Strings(resultArr)
  strInInt, _ := strconv.ParseInt(resultArr[len(resultArr)-1], 10, 64)
  return strInInt
}

func main() {
	fmt.Println(MaxRot(56789))
}