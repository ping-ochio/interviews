package inter

import (
	"fmt"
)

// --------------------- INTERFACE DECALRATION ---------------------------------
type User interface {
	Position() string
}

// --------------- STRUCTS THAT IMPLEMENTS USER INTERFACE ----------------------
type Admin struct {
	Name    string
	Surname string
}
type Employee struct {
	Name    string
	Surname string
}

// ----------- FUNCTIONS THAT IMPLEMENT THE USER INTERFACE ---------------------
func (e Admin) Position() string {
	return e.Name + " " + e.Surname + " is Administrator"
}
func (e Employee) Position() string {
	return e.Name + " " + e.Surname + " is Employee"
}

// ----------------- PRINT THE USER POSITION -----------------------------------
func Auth(user User) {
	fmt.Printf("| %s | \n", user.Position())
}
