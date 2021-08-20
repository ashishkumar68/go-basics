package main

import (
  "fmt"
  "net/http"
)

type hotdog int

func (val hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "This is a test example of handler")
}

func main() {
  var val hotdog
  http.ListenAndServe(":8080", val)
}
