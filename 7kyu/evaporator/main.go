package main 
import "fmt"

func Evaporator(content float64, evapPerDay int, threshold int) int {
  days := 0
  i := content;
  for i > (content * float64(threshold))/100 {
	i = i - (i * float64(evapPerDay)/100)	
    days++
  }
  return days
}

func main() {
	fmt.Println(Evaporator(10, 10, 10))
}