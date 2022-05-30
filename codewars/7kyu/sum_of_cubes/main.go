package main
import (
	"fmt"
	"math"
)

func SumCubes(n int) int {
  result := 0
  for i := 1; i <= n; i++ {
	  result += int(math.Pow(float64(i), float64(3)))
  }
  return result
}

func main() {
  fmt.Println(SumCubes(2))
}