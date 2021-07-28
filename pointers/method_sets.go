package main

import (
  "fmt"
)

type Person struct {
  firstName string
  lastName string
  age int
}

type VoteEligibleCheck interface {
  isVotingEligible() bool
}

func (person Person) isVotingEligible() bool {
  if (person.age >= 18) {
    return true
  } else {
    return false
  }
}

func (person Person) getFullName() string {
  return fmt.Sprintf("%v %v", person.firstName, person.lastName)
}

func (person *Person) setFirstName(firstName string) {
  person.firstName = firstName
}

func (person *Person) setLastName(lastName string) {
  person.lastName = lastName
}
func (person *Person) setAge(age int) {
  person.age = age
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
}
