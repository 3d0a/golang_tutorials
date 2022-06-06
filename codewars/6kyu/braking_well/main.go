package main

import (
	"fmt"
	"math"

)

func Dist(v, mu float64) float64 {                  // suppose reaction time is 1
	metersPerSec := v * 1000 / 3600
	return metersPerSec * metersPerSec / (2 * mu * 9.81) + metersPerSec
}

func Speed(d, mu float64) float64 {                 // suppose reaction time is 1
	speed         := (math.Sqrt(2 * mu * 9.81 * d))
	beforeBrake   := speed
	return (math.Sqrt(2 * mu * 9.81 * (d - beforeBrake))) * 3600 / 1000 
}

func main() {
	fmt.Println(Dist(144, 0.3))
	fmt.Println(Dist(92, 0.5))

  fmt.Println(Speed(159, 0.8))
  fmt.Println(Speed(164, 0.7))

}
