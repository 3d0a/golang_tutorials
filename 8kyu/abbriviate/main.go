package main
import (
	"fmt"
    "strings"
)

func AbbrevName(name string) string {
    array_of_abbr := make([]string, 0)
	for _, char := range strings.Split(name, " ") {
		array_of_abbr = append(array_of_abbr, strings.ToUpper(string(char[0])))
	}
	return strings.Join(array_of_abbr,".")
}

func main() {
    fmt.Println(AbbrevName("john kollins"))
}