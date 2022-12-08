/*
Here we make use of json.NewEncoder to literally pass the data directly to Stdout
(straight to the wire, that is, pure data to another application).
Once we have the data that we want to send, in this case in the slice people,
we generate the new encoder with "json.NewEncoder()" and the argument "os.Stdout",
that returns a pointer that we use with the function that it gives us the
JSON Encode package like this "pointer.Encode(whatever we want to send)",
for this particular case we pass the argument "people" which is a slice of "strutct person"
*/

package main

import (
	"encoding/json"
	"os"
)

type person struct {
	Name    string
	Surname string
	Age     int
}

func main() {

	p1 := person{
		Name:    "Peter",
		Surname: "Parker",
		Age:     20,
	}

	p2 := person{
		Name:    "Iron",
		Surname: "Man",
		Age:     45,
	}
	people := []person{p1, p2}

	point := json.NewEncoder(os.Stdout)
	point.Encode(people)
}
