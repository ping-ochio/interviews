package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	_, e := os.Create("test.txt")
	if e != nil {
		fmt.Printf("error:%v", e)
	}
	//	fmt.Printf("%T", *archi)
	data := `[{"name": "Peter", "surname": "Parker", "age": 20},{"name": "Iron", "surname": "Man", "age": 45}]`
	bs := []byte(data)
	err := ioutil.WriteFile("test.txt", bs, 644)
	if err != nil {

		log.Fatal(err)
	}
}
