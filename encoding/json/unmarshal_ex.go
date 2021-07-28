package main

import (
  "fmt"
  "encoding/json"
)

type Person struct {
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Age int `json:"age",omitempty`
}

func main() {
  serialisedBlob := []byte(`
    [
    {"first_name":"John","last_name":"Smith","age":23},
    {"first_name":"Kevin","last_name":"Smith","age":17},
    {"first_name":"John","last_name":"Doe"}
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
