package main
import (
	"fmt"
)

func search(nums []int, target int) int {
	return binary(nums, 0, len(nums) - 1, target)
}

func binary(nums []int, left, right, target int) int {
	if right >= left {
		mid := (left + right)/2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			return binary(nums, left, mid - 1, target)
		} else if target > nums[mid] {
			return binary(nums, mid + 1, right, target)
		}
	}
	return -1
}

func main() {
	nums := []int{-1,0,3,5,9,12}
	fmt.Println(search(nums, 12))
	
}
