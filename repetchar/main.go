 // Counts the number of times a character is repeated given a string

package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "oijloiurwqlkoooooKJÑLASOoij"
	var ch byte = 'l'

	cont := 0
	for i := 0; i < len(str); i++ {
		if ch == str[i] {
			cont++
		}
	}
	fmt.Printf("The number of times %q is repeted is %d\n", ch, cont)
	
	//-----------------------------------------------------------------
	// Same result using the Count() function from the strings package.
	chi := "l"
	str1 := strings.Count(str, chi)
	fmt.Println(str1)
}
