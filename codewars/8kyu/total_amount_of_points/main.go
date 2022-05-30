package main
import (
	"fmt"
	"strings"
)

func Ponits(games []string) int {
  resultPoints := 0
  for _, results := range games {
    winLose := strings.Split(results, ":")
    if []rune(winLose[0])[0] > []rune(winLose[1])[0] {
		resultPoints +=3	
	} else if []rune(winLose[0])[0] < []rune(winLose[1])[0] {
		continue
	} else {
		resultPoints +=1
	}
  }
  return resultPoints
}

func main() {
	fmt.Println(Ponits([]string{"2:2","3:4"}))
}