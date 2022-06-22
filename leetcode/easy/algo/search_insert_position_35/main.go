package main
import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	return binary(nums, 0, len(nums) - 1,target)
}

func binary(nums []int, left, right, target int) int{
  var mid int = (left + right) / 2
  if left >= right && nums[mid] >= target {
    return mid
  }
  if left >= right {
    return mid + 1
  }
  if nums[mid] == target {
    return mid
  }
  if nums[mid] > target {
    return binary(nums, left, mid - 1, target)
  } else if nums[mid] < target {
    return binary(nums, mid + 1, right, target)
  }
  return -1
}

func main() {
	arr := []int{1,3}
	fmt.Println(searchInsert(arr, 0))
}
