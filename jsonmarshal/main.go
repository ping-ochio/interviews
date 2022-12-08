/*
Working with JSON using JSON Marshal and JSON Unmarshal
and tags in stuct so we can use custom inputs name from json data
eg.: "name -> Name"
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//--------------------------------------------------------------------------------

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

//--------------------------------------------------------------------------------

func main() {
	c := []string{"hi", "there"}
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error:", err)

	}
	log.Printf("%q", b)
	os.Stdout.Write(b)

	n := []string{}

	err1 := json.Unmarshal(b, &n)
	if err1 != nil {
		fmt.Println("error:", err1)

	}
	fmt.Printf("\n%+v", n)

	//--- Now implementing person struct --------------------------------------------

	fmt.Println("\n\n---------- WORKING WITH STRUCTS TYPE ---------")
	data := `[{"name": "Peter", "surname": "Parker", "age": 20},{"name": "Iron", "surname": "Man", "age": 45}]`

	// Unmarshal receives "[]byte" and a pointer to target output
	// in this case "people"
	bs := []byte(data)

	fmt.Println("\n- PRINTING data\n", data)
	var people []Person //Slice of "Person" type where to point the unmarsalled data

	err2 := json.Unmarshal(bs, &people)

	if err2 != nil {
		fmt.Println("error:", err2)

	}
	fmt.Println("\n- PRINTING UNMARSHALLED DATA\n", people)

	fmt.Println("\n- PRINTING CLEAN DATA ")
	for _, per := range people {
		fmt.Printf("Name: %s %s and he has %d years old\n", per.Name, per.Surname, per.Age)
	}
}
