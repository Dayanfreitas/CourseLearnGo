package main

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{10, 11, 12}, 33},
		test{[]int{-5, 0, 5, 10}, 10},
	}

	for _, v := range tests {
		sum := soma(v.data...)
		if sum != v.answer {
			t.Error("Expected:", v.answer, "Got", sum)
		}
	}
}

func TestSoma(t *testing.T) {
	x := soma(3, 2, 1)
	r := 6
	if x != r {
		t.Error("Expected:", r, "Got", x)
	}
}

func ExempleSoma(t *testing.T) {
	fmt.Println(soma(3, 2, 1))
	//output 6
}

// func TestSoma2(t *testing.T) {
// 	x := soma(3, 2, 1)
// 	r := 7

// 	if x != r {
// 		t.Error("Expected:", r, "Got", x)
// 	}
// }

func TestMultplicação(t *testing.T) {
	x := multiplica(10, 10)
	r := 100

	if x != r {
		t.Error("Expected:", r, "Got", x)
	}
}
