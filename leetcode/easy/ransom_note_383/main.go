package main
import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
  for _, char := range ransomNote {
		if strings.Count(ransomNote, string(char)) > strings.Count(magazine, string(char)) {
			return false
		}
	}
	return true
}

func main() {

}
