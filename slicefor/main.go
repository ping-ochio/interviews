package main

import "fmt"

func main() {
	val := [][]string{
		{"Batman", "Boss", "dressed in black"},
		{"Robin", "Assistant", "colored dress"},
	}

	for pos := range val {
		fmt.Printf("Record nº: %d %s\n", pos, val[pos][:])
	}

	for p := range val {
		fmt.Printf("Record %d\n", p)

		for i, v := range val[p][:] {
			fmt.Printf("\tIndex: %d \tValue: %s\n", i, v)
		}
	}
}
