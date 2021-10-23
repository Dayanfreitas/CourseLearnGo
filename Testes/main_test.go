package main

import "testing"

func TestSoma(t *testing.T) int {
	x := Soma(3, 2, 1)
	r := 6
	if x != 6 {
		t.Error("Expected:", x, "Got", teste)
	}

}
