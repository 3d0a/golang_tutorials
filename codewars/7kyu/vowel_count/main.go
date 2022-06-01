package main
import (
	"fmt"
)

func GetCount(str string) (count int) {
	vowelMap := make(map[rune]int)
	vowelMap['a'] = 0
	vowelMap['e'] = 0
	vowelMap['i'] = 0
	vowelMap['o'] = 0
	vowelMap['u'] = 0
	for _, char := range str {
		if _, ok := vowelMap[char]; !ok {
			continue
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(GetCount("abracadabra"))
}
