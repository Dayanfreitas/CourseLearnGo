package main

import (
	"encoding/json"
	"fmt"
)

//UPPERCASE -> PUBLIC
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Profissao string
	Conta     float64
}

func main() {
	james_bonde := Pessoa{
		Nome:      "James",
		Sobrenome: "Bond",
		Idade:     40,
		Profissao: "Agente secreto",
		Conta:     10000.00,
	}
	darthvader := Pessoa{
		Nome:      "Darth",
		Sobrenome: "Vader",
		Idade:     50,
		Profissao: "Vilão intergalático",
		Conta:     500000.43,
	}

	j, err := json.Marshal(james_bonde)

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))
}
