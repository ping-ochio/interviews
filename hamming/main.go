
// Calculates the hamming distance given two strings of the same length

package main

import "fmt"

func main() {

	var text1 string = "patitosw"
	var text2 string = "paratosa"
	var distance int = 0

	if len(text1) != len(text2) {
		fmt.Println("ERROR! Strings must be the same length")
	}

	for i := 0; i < len(text1); i++ {
		if text1[i] != text2[i] {
			distance++
		}
	}
	fmt.Printf("The hamming distance is %d\n", distance)
}
