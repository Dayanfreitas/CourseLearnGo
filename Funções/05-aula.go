package main

import "fmt"

type pessoa struct {
	name      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	especializacao string
	salario        float32
}

type arquiteto struct {
	pessoa
	tipoDeConstrucao string
}

func (x arquiteto) hello() {
	fmt.Println("Meu nome é", x.name, "bom dia e eu sou arquiteto")
}

func (x dentista) hello() {
	fmt.Println("Meu nome é", x.name, "bom dia sou dentista")
}

type gente interface {
	hello()
}

func serhumano(g gente) {
	g.hello()
}
func main() {
	pessoa1 := pessoa{name: "mauricio", sobrenome: "Da Silva", idade: 50}
	pessoa2 := pessoa{name: "João", sobrenome: "Feitosa", idade: 43}

	mr_dentes := dentista{pessoa1, "cirurgião", 10000}
	mr_predio := arquiteto{pessoa2, "civil"}

	// mr_dentes.hello()
	// mr_predio.hello()

	serhumano(mr_dentes)
	serhumano(mr_predio)
}
