package main

import (
	"fmt"
)

func Divisors(n int ) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 2
	for i := 2; i <= n/2 ; i++{
		if  n%i == 0 {
			result += 1
		}
	}
	return result
}

func main() {
	fmt.Println(Divisors(10))
}
