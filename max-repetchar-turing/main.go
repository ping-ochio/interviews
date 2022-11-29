package main

import (
	"fmt"
	"regexp"
)

func main() {

	var char string
	max := 0
	s := "abcddefda1111133333333"
	r := regexp.MustCompile(`[a-z]`)
	a := r.FindAllString(s, -1)

	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		count := 0
		val := a[i]
		for j := 0; j < len(a); j++ {
			if val == a[j] && val != "" {
			
			        // It already counted this letter, 
			        // I delete it so that it doesn't count it again
				a[j] = "" // It already counted this letter, I delete it so that it doesn't count it again
				count++
			}
		}

		if count > max {
			max = count
			char = val
		}
	}

	fmt.Println(char, max)

}
