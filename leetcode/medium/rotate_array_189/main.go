package main
import (
	"fmt"
)

func rotate(nums []int, k int) []int {
	resultArr := []int{}
	if k >= len(nums) {k = k % len(nums)}
	resultArr = append(nums[len(nums) - k:], nums[:len(nums) - k]...)
	return resultArr
}

func main() {
	nums := []int{1,2,3,4,5,6,7}
	fmt.Println(rotate(nums, 3))
}
