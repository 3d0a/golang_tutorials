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
  for key, value := range mapCount {
    var tupple Tuple
	tupple.Char, tupple.Count = key, value
	finalArr = append(finalArr, tupple)
  }
  return finalArr
}

func main() {
  fmt.Println(OrderedCount("abracadabra"))
}