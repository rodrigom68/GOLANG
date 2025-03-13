package main

import "fmt"

type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	rua    string
	numero int
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "João"
	u.idade = 20
	fmt.Println(u)

	enderecoExemplo := endereco{rua: "Rua dos Bobos", numero: 0}

	usuario2 := usuario{nome: "João", idade: 20, endereco: enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 20}
	fmt.Println(usuario3)
}
