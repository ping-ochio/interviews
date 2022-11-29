package main

import (
	v "interfaces/inter"
)

// --------------------- MAIN FUNCTION -----------------------------------------
func main() {
	/*var new_employee v.Admin
	new_employee.Name = os.Args[1]
	new_employee.Surname = os.Args[2]*/

	users := []v.User{
		v.Admin{Name: "Jhon", Surname: "Wick"},
		v.Admin{Name: "Peter", Surname: "Cetera"},
		v.Employee{Name: "Hannibal", Surname: "Lecter"},
		v.Employee{Name: "Jack", Surname: "Sparrow"},
	}

	for _, pos := range users {

		v.Auth(pos)
	}
	
	//v.Auth(new_employee)

}

//-------------------------------------------------------------------------------
