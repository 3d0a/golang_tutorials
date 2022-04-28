package main
import (
	"fmt"
	"math"
)

func SquareSum(numbers []int) int {
	result := 0
    for _, number := range numbers {
        result += int(math.Pow(float64(number),2))
	}
	return result
}

func main() {
	array := []int{0,3,4,5}
    fmt.Println(SquareSum(array))
}