package main

import (
  "fmt"
)

type Person struct {
  FirstName string
  LastName string
  Age int
}

func (p *Person) Speak() {
  fmt.Println("My name is", p.FirstName, p.LastName, ".")
}

type Human interface {
  Speak()
}

func saySomething(h Human) {
  h.Speak()
}

func main() {
  john := Person{
    FirstName: "John",
    LastName: "Doe",
    Age: 33,
  }

  // this call fails, uncomment and run to test.
  //saySomething(john)

  // this one works
  saySomething(&john)
  // and this one also works, appearently we can't use non-pointer(T)
  // to call a receiver function with pointer(*T) through interface
  // but we can directly using the variable
  john.Speak()
}
