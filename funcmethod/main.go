/*
Here we used method-function associated to a strtct type

*/

package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func (r person) presentar() {
	fmt.Println("Hi my name is", r.name, r.surname, "i'm", r.age)
}

func main() {
	p := person{
		name:    "Karlota",
		surname: "Gimains",
		age:     25,
	}
	p1 := person{
		name:    "Simon",
		surname: "Templar",
		age:     36,
	}

	p.presentar()
	p1.presentar()

}
