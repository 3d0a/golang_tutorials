package main
import (
	"strings"
)

func isAnagram(s string, t string) bool {
  for _, char := range s  {
		if strings.Count(s, string(char)) != strings.Count(t, string(char)) {
			return false
		}
	}
	return true
}

func main() {

}
