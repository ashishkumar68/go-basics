package main

import (
    "fmt"
)

type Person struct {
    fName string
    lName string
}

func (p *Person) printName() {
    fmt.Println(p.fName, p.lName)
}

type human interface {
    printName()
}

func saySomething(h human) {
    h.printName()
}

func main() {
    psn := Person{
        fName: "James",
        lName: "Bond",
    }

    psn2 := &Person{
        fName: "Captain",
        lName: "Cool",
    }

    // The following will throw error if the variable is not Pointer type
    //saySomething(psn)
    saySomething(&psn)

    saySomething(psn2)
    //fmt.Println(psn.printName())
}
