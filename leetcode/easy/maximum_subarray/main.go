package main
import (
	"fmt"
	"sort"
)

func maxSubArray(nums []int) int {
	currentSum := 0
	maxSum := 0
	elemsLessZero := true

	for _, num := range nums {
		if num >= 0 {elemsLessZero = false}
		if currentSum + num > 0 {
			currentSum += num
		} else {
			currentSum = 0
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	if elemsLessZero {
		sort.Ints(nums)
		return nums[len(nums) - 1]
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2,-3}))
}
