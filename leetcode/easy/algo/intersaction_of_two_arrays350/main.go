package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	resultArr := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				resultArr = append(resultArr, nums1[i])
				if j != len(nums2)-1 {
					nums2 = append(nums2[:j], nums2[j+1:]...)
				} else {
					nums2 = append(nums2[:j], []int{}...)
				}
				break
			}
		}
	}
	return resultArr
}

func main() {
	arr1 := []int{9, 4, 9, 8, 4}
	arr2 := []int{4, 9, 5}
	fmt.Println(intersect(arr1, arr2))
}
