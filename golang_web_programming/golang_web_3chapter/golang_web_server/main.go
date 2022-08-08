package main
import (
	"fmt"
	"net/http"
)


func helloHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func worldHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/world", worldHandler)
	server.ListenAndServe()
}
