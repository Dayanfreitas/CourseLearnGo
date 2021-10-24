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
	sb := []byte(`[
		{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente secreto","Conta":10000},
		{"Nome":"Darth","Sobrenome":"Vader","Idade":50,"Profissao":"Vilão intergalático","Conta":500000.43}
	]`)
	var persons []Pessoa

	err := json.Unmarshal(sb, &persons)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(persons)

}
