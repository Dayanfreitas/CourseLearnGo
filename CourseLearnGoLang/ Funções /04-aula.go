package main

import "fmt"

type pessoa struct {
	name  string
	idade int
}

func (p pessoa) hello() {
	fmt.Println(p.name, "diz bom dia")
}

func main() {
	mauricio := pessoa{name: "mauricio", idade: 18}
	fmt.Println(mauricio)
	mauricio.hello()
}
