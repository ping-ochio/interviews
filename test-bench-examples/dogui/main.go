// Package dogui is a simulation to claculate a dog real age :)
package dogui

// Age convert dog age to human age
func Age(n int) int {
	return n * 7

}

// "AgeTwo" the same "Age" but more slow(used to test benchmarks)
func AgeTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
