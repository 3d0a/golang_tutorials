package main

import (
  "fmt"
  "net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Hello world, %s!", request.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", Handler)
  http.ListenAndServe(":8080", nil)
}