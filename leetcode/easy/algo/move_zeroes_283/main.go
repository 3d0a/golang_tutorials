package main
import (
	"fmt"
)

func moveZeroes(nums []int)  {
	zeros := 0
	for fast := 0; fast < len(nums); fast++  {
		if nums[fast] != 0 && nums[zeros] == 0 {
			nums[fast], nums[zeros] = nums[zeros], nums[fast]
		}
		if nums[zeros] != 0 {
			zeros++
		}
	}
}

func main() {
	arr := []int{0,3,0,4,0,5,13,0}
	fmt.Println(arr)
	moveZeroes(arr)
	fmt.Println(arr)
}
