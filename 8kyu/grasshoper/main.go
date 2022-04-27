package main
import "fmt"

func combat(health, damage float64) float64 {
    switch {
	case damage > health:
	    return 0
	default:
	    return health - damage
	}

}

func main() {
    fmt.Println(combat(4,6))
}