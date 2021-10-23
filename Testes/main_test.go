package main

import "testing"

func TestSoma(t *testing.T) {
	x := soma(3, 2, 1)
	r := 6
	if x != r {
		t.Error("Expected:", r, "Got", x)
	}
}

func TestSoma2(t *testing.T) {
	x := soma(3, 2, 1)
	r := 7

	if x != r {
		t.Error("Expected:", r, "Got", x)
	}
}

func TestMultplicação(t *testing.T) {
	x := multiplica(10, 10)
	r := 100

	if x != r {
		t.Error("Expected:", r, "Got", x)
	}
}
