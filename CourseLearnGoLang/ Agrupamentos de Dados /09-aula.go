package main

import "fmt"

func main() {
	amigos := map[string]int{
		"alfredo": 423423423,
		"joana":   242343,
	}

	if v, ok := amigos["fantasma"]; !ok {
		fmt.Println("Não existe", v)
	}

	fmt.Println(amigos["joana"])
}
