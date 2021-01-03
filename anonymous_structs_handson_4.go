package main

import (
	"fmt"
)

func main() {
	john := struct {
		// bool phone number representing active or inactive phone number
		phonenumbers map[string]bool
		addresses    []string
		name         string
	}{
		phonenumbers: map[string]bool{
			"8727373274": true,
			"8264637734": false,
		},
		addresses: []string{
			"test address1", "test address 2",
		},
		name: "John Smith",
	}

	fmt.Println(john.name)
	for number, isActive := range john.phonenumbers {
		if isActive {
			fmt.Printf("%s is Active", number)
		} else {
			fmt.Printf("%s is InActive", number)
		}
		fmt.Println()
	}
	for _, address := range john.addresses {
		fmt.Println(address)
	}
}
