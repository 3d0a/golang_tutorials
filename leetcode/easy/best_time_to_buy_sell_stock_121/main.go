package main
import (
	"fmt"
)

func maxProfit(prices []int) int {
	maxIncome  := 0
	currentMax := prices[len(prices)-1]
	for i := len(prices) - 1; i >= 0; i-- {
		if currentMax - prices[i] > maxIncome {
			maxIncome = currentMax - prices[i]
		}
		if currentMax < prices[i] {
			currentMax = prices[i]
		}
	}
	return maxIncome
}



func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
