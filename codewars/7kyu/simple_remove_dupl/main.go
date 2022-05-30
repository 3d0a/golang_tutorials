package main
import (
  "fmt"
)

func Solve(arr []int) []int {
  numberMap     := make(map[int]bool)
  semiResultArr := make([]int,0)
  resultArr     := make([]int,0)
  for i := len(arr) - 1; i>=0; i-- {
    if _, value := numberMap[arr[i]]; !value {
		numberMap[arr[i]] = true
		semiResultArr = append(semiResultArr, arr[i])
	}
  }
  for i := len(semiResultArr) - 1; i>=0; i-- {
    resultArr = append(resultArr, semiResultArr[i])
  } 
  return resultArr
}

func main() {
  fmt.Println(Solve([] int{3,4,4,3,6,3}))
}