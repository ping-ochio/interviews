package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	pass := `5up3r53c37p455`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), 4)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(pass)
	fmt.Println(string(bs))
	//wrongpass := `wrongpass`
	err = bcrypt.CompareHashAndPassword(bs, []byte(pass))
	if err != nil {
		fmt.Println("WRONG PASSWORD, ACCES DENIED..!!")
		return
	}
	fmt.Println("ACCESS SUCCESSFULLY ")
}
