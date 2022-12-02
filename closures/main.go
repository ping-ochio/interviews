/*
CLOSURES
We made of anonymous functions to enclose the scope of a variable.
When declaring a variable to which an anonymous function is assigned by calling a function that returns an anonymous function, that variable is of type function and since it is an anonymous function,
it also receives the environment of the scope of that function where the function is enclosed. initial variable.
*/
package main

import "fmt"

func main() {
	//---- "y()" contains its own scope of "x", so every time we
	// make with the same identifier "j = y()" increment only j value
	y := foo()
	j := y()
	fmt.Println("j value:", j)
	j = y()
	fmt.Println("j value:", j)
	//---- 	new scope ---
	u := foo()
	h := u()
	fmt.Println("h value", h)
	h = u()
	fmt.Println("h value:", h)

}

func foo() func() int {

	var x int
	return func() int {

		x++
		return x

	}
}
