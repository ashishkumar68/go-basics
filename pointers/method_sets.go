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

type VoteEligibleCheck interface {
  isVotingEligible() bool
}

func (person Person) isVotingEligible() bool {
  if (person.Age >= 18) {
    return true
  } else {
    return false
  }
}

func (person Person) getFullName() string {
  return fmt.Sprintf("%v %v", person.FirstName, person.LastName)
}

func (person *Person) setFirstName(firstName string) {
  person.FirstName = firstName
}

func (person *Person) setLastName(lastName string) {
  person.LastName = lastName
}
func (person *Person) setAge(age int) {
  person.Age = age
}

func main() {
  john := Person{}
  fmt.Println(john)
  fmt.Println(&john)
  fmt.Printf("%T \n", &john)

  john.setFirstName("John")
  john.setLastName("Smith")
  john.setAge(22)

  fmt.Println(john)
  fmt.Println(john.getFullName())
  fmt.Println(john.isVotingEligible())

  serialised, err := json.Marshal(john)
  if err != nil {
    fmt.Println("error occured while serializing john:", err)
  } else {
    fmt.Println(string(serialised))
  }
}
