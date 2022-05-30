package main
import (
	"fmt"
	"sort"
)

func SortNumbers(number []int) []int {
  resultArray := number
  if number != nil {
	sort.Ints(resultArray)
	return resultArray
  }
  return nil
}

func main() {
	fmt.Println(SortNumbers([]int{}))
}
