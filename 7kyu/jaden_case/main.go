package main
import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
    result_string := make([]string, 0)
	for _, word := range strings.Split(str, " ") {
		result_string = append(result_string, strings.Title(word))
	}
	return strings.Join(result_string, " ")
}

func main() {
    fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
}