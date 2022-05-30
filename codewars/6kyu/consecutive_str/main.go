package main
import (
  "fmt"
)

func LongestConsec(strarr []string, k int) string {
  longestString := ""
  semiString    := ""
  if len(strarr) ==0 || k > len(strarr) || k <= 0 {
    return ""
  }
  for i := 0; i < len(strarr) - k + 1; i++ {
	semiString = ""
    for j := i; j <= i + k -1; j++ {
	  semiString +=  strarr[j]
	}
	if len(semiString) > len(longestString) {
	  longestString = semiString
	}
  }
  return longestString
}

func main() {
  fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 5))
}