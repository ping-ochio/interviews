package main

import (
	"fmt"
	"test/dogui"
)

type dog struct {
	name string
	age  int
}

func main() {
	fido := dog{
		name: "Fido",
		age:  dogui.Age(10),
	}
	fmt.Println(fido)
	fmt.Println(dogui.AgeTwo(20))

}
