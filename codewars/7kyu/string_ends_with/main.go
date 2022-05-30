package main
import "fmt"

func StringEnds(str, ending string) bool {
  if len(ending) == 0 {
	  return true
  } else if  len(str) == 0 {
	  return false
  }
  i := len(str) - 1 
  for j := len(ending) - 1; j >=0; j-- {
	if ending[j] == str[i] && i != -1 {
	  i--
	  continue
	} else {
		return false
	}
  }
  return true
} 

func main() {
  fmt.Println(StringEnds(" ", ""))
}