package main
import ( 
	"fmt"
	"strings"
)

func DNAtoRNA(dna string) string{
	array := make([]string, 0)
	for i, char := range dna {
		if string(char) == "T" {
			array = append(array, "U")
		} else {
			array = append(array, string(char))
		}
        _ = i
	}
	return strings.Join(array, "")
}

func main() {
    fmt.Println(DNAtoRNA("helloT"))
}