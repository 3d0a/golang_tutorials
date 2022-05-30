package main
import ( 
  "fmt"
  "strings"
)

func Capitalize(st string, arr []int) string {
  arrOfStrings := strings.Split(st, "")
  for _, index := range arr {
	  if index >= len(st) {
		continue
	  } else {
        arrOfStrings[index] = strings.ToUpper(arrOfStrings[index])
	  }
  }
  return strings.Join(arrOfStrings, "")
}

func main() {
  fmt.Println(Capitalize("abcdef", []int {1,2,5}))
}
