package main
import (
	"fmt"
	"strings"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {return false}
  return permutation(strings.Split(s1, ""), 0, s2)
}

func permutation(arr []string, left int, s2 string) bool {
  hasPermutation := false
  if left == len(arr) {
      return strings.Contains(s2, strings.Join(arr, ""))
  }
  for i := left; i < len(arr); i++ {
    arr[left], arr[i] = arr[i], arr[left]
    hasPermutation = permutation(arr, left + 1, s2)
    if hasPermutation == true {break}
    arr[left], arr[i] = arr[i], arr[left]
  }
  return hasPermutation
}

func main() {
	str2 := "dinitrophenylhydrazine"
	str :=  "acetylphenylhydrazine"
	fmt.Println(checkInclusion(str,str2))
}
