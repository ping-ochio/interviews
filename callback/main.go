package main

import "fmt"

// --------------------------------------------
func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstSum := sum(list...)
	fmt.Println("The total of the sum is:", firstSum)

	totalPairs := sumpairs(sum, list...) //Only "sum" without "()" since it is only an argument
	fmt.Println("The total of the sum of pairs is:", totalPairs)
}

// --------------- TOTAL SUM ------------------
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// ------------- TOTAL PAIRS SUM --------------
func sumpairs(n func(x ...int) int, vi ...int) int { // Here and in this case "n" is "sum"
	storePairs := []int{}
	for _, v := range vi {

		if v%2 == 0 {
			storePairs = append(storePairs, v)
		}
	}
	return n(storePairs...) //n can be sum

}

//--------------------------------------------
