package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i, elem := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target - elem {return []int{i, j}}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3,2,4}, 7))
}
