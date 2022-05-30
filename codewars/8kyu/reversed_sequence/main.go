package main 
import "fmt"

func ReversedSeq(n int) []int {
    array := make([]int, 0)
	for n > 0 {
		array = append(array, n)
		n--
	}
	return array
}

func main() {
    fmt.Println(ReversedSeq(4))
}