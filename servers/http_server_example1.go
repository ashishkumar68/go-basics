package main

import (
  "fmt"
  "log"
  "net/http"
)

type hotdog int

func (val hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  log.Fatalln(fmt.Fprintln(w, "This is a test example of handler"))
}

func main() {
  var val hotdog
  log.Fatalln(http.ListenAndServe(":8080", val))
}
