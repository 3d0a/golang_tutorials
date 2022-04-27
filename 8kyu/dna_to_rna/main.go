package main
import ( 
	"fmt"
	"reflect"
)

func DNAtoRNA(dna string) {
	for i, char := range dna {
        fmt.Println(reflect.TypeOf(char), reflect.TypeOf(i), string(char), reflect.TypeOf(string(char)))
	}
}

func main() {
    DNAtoRNA("hello")
}