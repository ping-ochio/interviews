package main

import "fmt"

func main() {

	p := func(s, l string) {
		fmt.Println(s, "say", l)
	}

	greet := p
	groot := p
	fmt.Printf("greet type: \"%T\"\n", greet)
	greet("greet", "Hi")
	groot("groot", "Bye")
}
