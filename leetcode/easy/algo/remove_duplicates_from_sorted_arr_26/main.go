package main
import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	mapDups := make(map[int]bool)
	dupCount := 0
	for i, elem := range nums {
		if ok, _ := mapDups[elem]; !ok {
			mapDups[elem] = true
		} else {
			nums[i] = -1000
			dupCount++
		}
	}
	slow := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] != -1000 && nums[slow] == -1000 {
			nums[i], nums[slow] = nums[slow], nums[i]
		}
		if nums[slow] != -1000 {
			slow++
		}
	} 
	return len(nums) - dupCount
}

func main() {
	arr := []int{0,0,0,1,2,3,3,3,4,5,6,6,7,7,7,9,10}
	removeDuplicates(arr)
	fmt.Println(arr)
}
