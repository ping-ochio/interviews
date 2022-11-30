package main

import "fmt"

// -------------- STRUCTS ------------------------
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle    // Promoted field
	fourwheels bool
}
type Sedan struct {
	vehicle // Promoted field
	luxury  bool
}

// --------------- MAIN ------------------------
func main() {

	v1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourwheels: false,
	}
	v2 := Sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Printf("The truck has %d doors, it is %s and it is %v that has four wheels\n", v1.doors, v1.color, v1.fourwheels)
	fmt.Printf("The sedan has %d doors, it is %s and it's %v that it's luxurious\n", v2.doors, v2.color, v2.luxury)

}
