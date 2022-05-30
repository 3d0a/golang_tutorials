package main
import (
	"fmt"
)

func main() {
	fmt.Println(FindOdd([]int{1,2,2,1,1,3,3}))
}

func  FindOdd(seq []int) int {
	repeatMap := make(map[int]int)
	returnVal := 0
	for _, digit := range seq {
		if _, ok := repeatMap[digit]; !ok {
			repeatMap[digit] = 1
		} else {
			repeatMap[digit] += 1
		}
	}
	for key, value := range repeatMap {
		if value % 2 != 0 {
			returnVal = key
			break
		}
	}
	return returnVal
}
