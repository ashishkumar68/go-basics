package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		fourWheel: true,
	}

	car1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Dark Blue",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(truck1.doors)
	fmt.Println(truck1.fourWheel)

	fmt.Println(car1)
	fmt.Println(car1.doors)
	fmt.Println(car1.luxury)
}
