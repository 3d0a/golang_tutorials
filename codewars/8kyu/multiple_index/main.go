package main
import "fmt"

func multipleOfIndex(ints []int) []int {
  resultArray := make([]int, 0)
  for index, value := range ints {
     if index != 0 && value % index == 0  {
		 resultArray = append(resultArray, value)
	 }
  }
  return resultArray
}

func main() {
  fmt.Println(multipleOfIndex([]int{ 22, -6, 32, 82, 9, 25}))
  for index, value := range []int{ 22, -6, 32, 82, 9, 25} {
	  fmt.Println(index, value)
  }
}