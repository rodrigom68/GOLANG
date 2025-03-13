package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Arquivo de herança")

	p1 := pessoa{"João", "Pedro", 20, 1.75}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.faculdade)
	fmt.Println(e1.idade)
	fmt.Println(e1.altura)
}
