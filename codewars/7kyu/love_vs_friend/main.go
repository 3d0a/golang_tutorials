package main
import "fmt"

func WordsToMarks(s string) int {
  result := 0
  for _, char := range s {
	result += int(char - 96)
  }
  return result
}

func main(){
  fmt.Println(WordsToMarks("love"))
}