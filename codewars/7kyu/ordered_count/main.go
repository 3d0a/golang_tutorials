package main
import (
  "fmt"
)

type Tuple struct {
  Char  rune
  Count int
}

func OrderedCount(text string) []Tuple {
  mapCount := make(map[rune]int)
  finalArr := make([]Tuple, 0)
  for _, char := range []rune(text) {
	if _, ok := mapCount[char]; !ok {
	  mapCount[char] = 1
	} else {
	  mapCount[char] += 1
	}
  }
  for _, char := range []rune(text) {
	if _, ok := mapCount[char]; !ok {
	  continue
	}
	var tuple Tuple
    tuple.Char, tuple.Count = char, mapCount[char]
	finalArr = append(finalArr, tuple)
	delete(mapCount, char)
  }
  return finalArr
}

func main() {
  fmt.Println(OrderedCount("abracadabra"))
}