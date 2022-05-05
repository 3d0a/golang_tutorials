package main
import (
  "fmt"
  "sort"
  "strings"
)

func TwoSort(arr []string) string {
  sort.Strings(arr)
  charArr := make([]string, 0)
  for _, char := range []rune(arr[0]) {
    charArr = append(charArr, string(char))
  }
  return strings.Join(charArr,"***")
}

func main() {
  fmt.Println(TwoSort([]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}))
}
