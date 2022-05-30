package main

import (
	"fmt"
)

func  MaxBall(v0 int) int {
  max_height := 0.0
  v := float64(v0) * 1000/3600
	for i := 0; ; i++ {
		h := v * float64(i)/10 - 0.5*9.81*float64(i*i)/100
		if h >= max_height {
			max_height = h
		} else {
			return i-1
		}
	}
}

func main() {
	fmt.Println(MaxBall(15))
}
