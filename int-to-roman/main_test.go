package main

import "testing"

func TestToroman(t *testing.T) {
	res := Toroman("56")
	if res != "LVI" {

		t.Errorf("Unsuccessful conversion, expected %s and it was obtained %s", "LVI", res)
	}

}

func TestSuma(t *testing.T) {
	sum := Suma(5, 3)
	if sum != 8 {

		t.Errorf("does not add up well, we wanted %d and it was obtained %d", 8, sum)
	}

}
