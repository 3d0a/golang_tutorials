package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	finalArr := []int{}
	i, j := 0, len(numbers)-1
	for i < len(numbers) && j >= 0 {
		if numbers[i]+numbers[j] == target {
			finalArr = append(finalArr, i+1, j+1)
			break
		} else if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		}
	}
	return finalArr
}

func main() {

}
