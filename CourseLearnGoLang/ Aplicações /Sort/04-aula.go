package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

func main() {
	cars := []carro{
		{nome: "Celta", potencia: 60, consumo: 20},
		{nome: "KA", potencia: 50, consumo: 14},
		{nome: "Porche", potencia: 1000, consumo: 5},
	}

	fmt.Println("INICIAL->", cars)

	fmt.Println("Potencia")
	sort.Slice(cars, func(i, j int) bool { return cars[i].potencia < cars[j].potencia })
	fmt.Println(cars)

	fmt.Println("Consumo")
	sort.Slice(cars, func(i, j int) bool { return cars[i].consumo > cars[j].consumo })
	fmt.Println(cars)

	fmt.Println("Lucro")
	sort.Slice(cars, func(i, j int) bool { return cars[i].consumo < cars[j].consumo })
	fmt.Println(cars)
	fmt.Println("_______")

}
