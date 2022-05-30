package main
import (
  "fmt"
  "sort"
  "strings"
)

func Solve(s string) bool {
  StringArr := strings.Split(s, "") 
  sort.Strings(StringArr)
  runeMap := make(map[rune] int)
  arrRune := []rune(strings.Join(StringArr, ""))
  if len(arrRune) == 0 {
    return false
  }
  for i := 0; i < len(arrRune); i++ {
	if _, ok := runeMap[arrRune[i]]; !ok {
	  runeMap[arrRune[i]] += 1
	} else {
	  return false
	}
    if (arrRune[i] >= 'a' && arrRune[i] <= 'z') {
	  if i != len(arrRune) - 1 && len(arrRune) > 1 && arrRune[i+1] != arrRune[i] + 1 {
        return false
	  } else if  len(arrRune) == 1 {
		return true
	  } 
	}
  }
  return true
}

func main() {
  fmt.Println(Solve("acb"))
}