package main
import "fmt"

func SmallestIntegerFinder(numbers []int) int {
    min := numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
    return min
}

func main() {
	array := []int{-333,4,-10000,5,777}
	fmt.Println(SmallestIntegerFinder(array))
}