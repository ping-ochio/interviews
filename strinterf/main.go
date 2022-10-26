package main

import "fmt"

var v string = `
{
	"Name":"Juan",
	"Age":"19",
	"Hobbies": 
	[
	"martial arts",
	"breakfast foods",
	"piano"
	]
}
`

//This declaration works like a global variable
//var values = []string{"valor1", "valor2"}

func main() {

	p := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	//This declaration of a slice only works inside as a local variable
	values := []string{"first", "second"}

	fmt.Printf("v is a type: %T\nvalues is a type: %T\n\n", v, values)
	//Using %q can see the "\n,\s,\t"
	fmt.Printf("%q\n", v)

	//Using %s print the v content whitout "\n,\s,\t"
	fmt.Printf("%s", v)

	for k, v := range p {
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)

		}

	}

}
