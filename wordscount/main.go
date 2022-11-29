//Only counts the words in the text, trims the empty spaces

package main

import (
	"fmt"
	"regexp"
	"strings"
)

// main function
func main() {
	str := "   this 	program	 count      only the words, not    the    spaces or tabs   "
	r := regexp.MustCompile(`\s+`)
	s := r.ReplaceAllString(str, " ")
	s2 := strings.TrimSpace(s)
	si := strings.Split(s2, " ")
	count := len(si)
	fmt.Printf("This text [... %s ...] \nhas %d words\n", str, count)
}
