//Reverses text using the for loop in two ways

package main

import "fmt"

func main() {

	str := "abcd"

	for pos := len(str) - 1; pos >= 0; pos-- {

		fmt.Printf("%c", str[pos])
	}
		fmt.Println("\n")
	
	//--------------------------------------------------
	
	str1 := ""
	fmt.Println("Here we can see how the for loop reorders the characters")
	for _, val := range str {
		str1 = string(val) + str1
		print(str1 + "-")
	}
	fmt.Println("\nFinal result:\n" + str1)

}
