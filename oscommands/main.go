package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

// with this simple program we can run system commands
// in this case only for linux
// makes a folder called "test" and subfolder called "folder"

//var prog = "setterm"
//var arg1 = "--cursor"
//var arg2 = "on"

var prog = "mkdir"
var arg1 = "-p"
var arg2 = "test/folder"

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		commands(prog, arg1, arg2)
	}

}

func commands(s string, arg1 string, arg2 string) {

	out, err := exec.Command(s, arg1, arg2).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)

}
