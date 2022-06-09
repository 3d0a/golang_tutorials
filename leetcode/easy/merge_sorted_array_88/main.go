package main
import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums = nums1[:m]
	for i := 0; i < n; i++ {
		nums = append(nums, nums2[i])
	}
	sort.Ints(nums)
}

func main() {

}
