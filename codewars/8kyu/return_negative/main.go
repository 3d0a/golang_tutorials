package main

import (
	"fmt"
	"math"
)

func MakeNegative(x int) int {
	isNegative := math.Signbit(float64(x))
    if isNegative {
		return x
	}
	return -1 * x
}

func main() {
    fmt.Println(MakeNegative(0))
}