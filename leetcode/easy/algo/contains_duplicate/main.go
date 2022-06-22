package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	mapCounter := make(map[int]int)
	for _, elem := range nums {
		mapCounter[elem]++
		if mapCounter[elem] > 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{3, 4, 5, 1, 1, 0}))
}
