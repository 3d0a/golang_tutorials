package main

import (
	"fmt"
)

func matrixReshape(mat [][]int, r int, c int) [][]int {
	resultMatrix := make([][]int, 0)
	for i := 0; i <= r*c; i++ {
		resultMatrix = append(resultMatrix)
	}
	return resultMatrix
}

func main() {
	fmt.Println(matrixReshape([][]int{{1,2}, {3,4}}, 1, 4))
}
