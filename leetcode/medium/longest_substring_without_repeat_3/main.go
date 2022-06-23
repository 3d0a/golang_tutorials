package main 
import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
  mapCountOfChars := make(map[string]int)
  stringArr := strings.Split(s, "")
  biggestSubstring := ""
  semiSubstring := ""
  for i := 0; i < len(stringArr); i++ {
      if _, ok := mapCountOfChars[stringArr[i]]; !ok {
        mapCountOfChars[stringArr[i]] = i
        semiSubstring += stringArr[i]
      } else {
        i = mapCountOfChars[stringArr[i]]
        mapCountOfChars = make(map[string]int)
        semiSubstring = ""
      }
      if len(semiSubstring) > len(biggestSubstring) {
        biggestSubstring = semiSubstring
      }
  }
  return len(biggestSubstring) 
}
