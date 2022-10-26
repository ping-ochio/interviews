
// This program counts how many numbers are inside the string "text"

package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "an1298jq???daº22121230001masA"
	re := regexp.MustCompile(`[0-9]`)
	str := re.FindAllString(text, -1)
	str1 := len(str)
	fmt.Printf("The entered text has %d searched characters\n", str1)
}
