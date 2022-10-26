package main

import "fmt"

/*
	Different ways to create a slice

	-- name := make{[]type, long, capacity}
	-- name := []type{elemet1, element2}
	-- var name = []type{elemet1, element2}
	-- Adding elements " name = append(name, element1, element2, ...)"
	-- Put or change a value in a position " name[position] = vlue "
*/

func main() {

	mode1()
	mode2()
	mode3()
}

//-------------------------------------------------------------------

func mode1() {

	fmt.Println("\n--- Here a slice of integer elements without size")
	Items := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("\tInitial slice: %v", Items)

	Items[2] = 5
	fmt.Printf("\n\n\tChange value at the second position: %v\n\n", Items)

	fmt.Println("\tPrint from position 0 to position 4")
	fmt.Printf("\t%v", Items[0:4])

}

//-------------------------------------------------------------------

func mode2() {

	fmt.Println("\n\n--- Here with size (Length and capacity)")
	values := make([]uint8, 0, 10)
	fmt.Printf("\tThe slice content: %v, length of slice: %d, and capacity %d", values, len(values), cap(values))

	fmt.Println("\n\n\tUsing the append method to add elements")
	values = append(values, 2, 55, 6)

	fmt.Printf("\n\tNow the slice content: %d, length of slice: %d, and capacity %d", values, len(values), cap(values))
}

//-------------------------------------------------------------------

func mode3() {

	fmt.Println("\n\n--- Here creating a slice using \"var values = []string{\"first\", \"second\"}\"  ")
	var values = []string{"first", "second"}
	fmt.Printf("\n\tSlice content: %v, length of slice: %d, and capacity %d\n\n", values, len(values), cap(values))

}

//-------------------------------------------------------------------
