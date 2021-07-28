package main

import (
  "fmt"
  "encoding/json"
)

type Person struct {
  FirstName string
  LastName string
  Age int
}

func main() {
  serialisedBlob := []byte(`
    [
    {"firstName":"John","lastName":"Smith","age":23},
    {"firstName":"Kevin","lastName":"Smith","age":17}
    ]
  `)
  persons := []Person{}
  err := json.Unmarshal(serialisedBlob, &persons)
  if err != nil {
    fmt.Println("Could not deserialized the persons data due to error:", err)
  } else {
    fmt.Println(persons)
  }
}
