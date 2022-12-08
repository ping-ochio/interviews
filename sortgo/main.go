package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name    string
	Surname string
	Age     int
	Sports  []string
}

func main() {

	inval := []int{25, 4, 3, 45, 40, 10, 1}
	strval := []string{"Paul", "Peter", "Mandy", "Mary"}

	fmt.Println(inval)
	sort.Ints(inval)
	fmt.Println(inval)
	fmt.Println(strval)
	sort.Strings(strval)
	fmt.Println(strval)

	// ------------------------------------------------------

	fmt.Printf("\n\n----------------------------- \n")

	u1 := user{Name: "Robert", Surname: "Miller", Age: 15, Sports: []string{"paddle", "football", "byking"}}
	u2 := user{Name: "Martin", Surname: "Bellis", Age: 33, Sports: []string{"swimming", "ride", "golf"}}
	u3 := user{Name: "John", Surname: "Hammond", Age: 34, Sports: []string{"basketball", "baseball", "bowling"}}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	fmt.Printf("\n----------- SORTING BY SURNAME AND SPORT ------------------ \n")
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	for _, val := range users {
		fmt.Printf("Name:%v %v %v\n", val.Name, val.Surname, val.Age)
		sort.Strings(val.Sports)
		for i, v := range val.Sports {

			fmt.Printf("\t%d %v\n", i+1, v)
		}

	}
}
