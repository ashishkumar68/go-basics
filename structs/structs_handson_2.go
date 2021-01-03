package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	flavours  []string
}

func (p person) getFullName() string {
	return fmt.Sprintf("%s %s", p.firstname, p.lastname)
}

func (p person) getFavIcecreamFlavours() []string {
	return p.flavours
}

func main() {
	john := person{
		firstname: "John",
		lastname:  "Smith",
		flavours: []string{
			"Vanilla",
			"Chocolate",
			"Hot Chocolate",
			"rum and coke",
		},
	}

	james := person{
		firstname: "James",
		lastname:  "Bond",
		flavours: []string{
			"Hot Fudge",
			"Vanilla",
		},
	}

	people := map[string]person{
		john.lastname:  john,
		james.lastname: james,
	}

	/** @var person per **/
	for lastname, per := range people {
		fmt.Println(per.getFullName(), "'"+lastname+"'")
		for key, fav := range per.getFavIcecreamFlavours() {
			fmt.Printf("Flavour %d %s for %s", (key + 1), fav, per.getFullName())
			fmt.Println()
		}
	}
}
