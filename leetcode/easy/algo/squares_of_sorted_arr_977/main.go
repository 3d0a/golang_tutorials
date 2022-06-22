package main
import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	finalArr := []int{}
	for _, elem := range nums {
		finalArr = append(finalArr, elem)
	}
	sort.Ints(finalArr)
	return finalArr
}

func main() {

}
