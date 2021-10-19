package main

import (
	"fmt"
)

type person struct {
	name string
}

func mudeMe(p *person) {
	(*p).name = "Dan"
}

func main() {
	person1 := person{name: "Dayan"}
	mudeMe(&person1)
	fmt.Println(person1)
}
