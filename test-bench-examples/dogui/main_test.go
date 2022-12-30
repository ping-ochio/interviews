package dogui

import (
	"fmt"
	"testing"
)

func TestAge(t *testing.T) {

	n := Age(10)
	if n != 70 {
		t.Error("Want", 70, "and got", n)
	}

}

func TestAgeTwo(t *testing.T) {

	n := AgeTwo(10)
	if n != 70 {
		t.Error("Want", 70, "and got", n)
	}

}

func ExampleAge() {
	fmt.Println(Age(10))
	//Output:
	//70
}
func ExampleAgeTwo() {
	fmt.Println(AgeTwo(10))
	//Output:
	//70
}
func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {

		Age(10)

	}

}
func BenchmarkAgeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {

		AgeTwo(10)

	}

}
