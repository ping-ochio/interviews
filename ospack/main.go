
// This program makes use of operating system options to check if a file exists given a path

package main

import (
	"fmt"
	"os"
)

func main() {

	if _, err := os.Stat("/home/user/.local/bin/file"); !os.IsNotExist(err) {
		fmt.Println("File not found")
	} else {

		fmt.Println("File exist")
	}

}
