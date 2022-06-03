package main

import (
	"fmt"
)

func HighestRank(nums []int) int {
	mapOfRepeats := make(map[int]int)
	max_value    := 0
	max_key      := 0
	for _, num := range nums {
		if _, ok := mapOfRepeats[num]; !ok {
			mapOfRepeats[num] = 1
		} else {
			mapOfRepeats[num] += 1
		}
	}
	for key, value := range mapOfRepeats {
		if value > max_value {
			max_value = value
			max_key   = key
		}	else if key > max_key && value == max_value {
			max_key = key
		}
	}
	return max_key
}

func main() {
	fmt.Println(HighestRank([]int{10, 8, 12, 7, 6, 4, 10, 12, 10 }))
}
