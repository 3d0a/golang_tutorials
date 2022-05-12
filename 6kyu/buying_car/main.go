package main
import (
  "fmt"
  "math"
)

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
  countMounths := 0
  leftMoney    := 0.0
  newPriceOld  := float64(startPriceOld)
  newPriceNew  := float64(startPriceNew)
  var resualtArr [2]int
  for leftMoney = 0; leftMoney + newPriceOld < newPriceNew; leftMoney += float64(savingperMonth)  {
    countMounths += 1
    if  countMounths % 2 == 0 {
      percentLossByMonth += 0.5
    }
	  newPriceNew = newPriceNew - (newPriceNew * percentLossByMonth)/100
    newPriceOld = newPriceOld - (newPriceOld * percentLossByMonth)/100
  }
  resualtArr[0], resualtArr[1] = countMounths, int(math.Round(leftMoney + newPriceOld - newPriceNew))
  return resualtArr
}

func main() {
  fmt.Println(NbMonths(2000, 8000, 1000, 1.5))
}