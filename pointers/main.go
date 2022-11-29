package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// --------------------------------------------
func main() {

	var x *int
	var p *int

	user := Person{"Peter", 25}

	content := 5

	x = &content

	fmt.Println(content)
	*x = 45
	fmt.Println(content)

	change(&content)

	fmt.Println(content)
	fmt.Println(user)
	change_age(&user)
	fmt.Println(user)

	// --------------------------------------------

	var v int = 19

	//Pointer p, reference the memory address of the variable v.
	p = &v

	fmt.Printf("\nV ariable v is: %d \n", v)

	fmt.Printf("Address of v is: %v \n", &v)

	fmt.Printf("The pointer p refers to the memory address: %v \n", p)

	fmt.Printf("When dereferencing the pointer we get the value: %d \n", *p)

}

// --------------------------------------------
func change(x *int) {
	*x = 67

}

// --------------------------------------------
func change_age(a *Person) {
	a.Age++
	fmt.Println("Username: ", a.Name)

}
