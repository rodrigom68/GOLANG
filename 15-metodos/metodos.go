package main //

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { // método salvar
	// método salvar
	// o método salvar é um método do tipo usuario
	fmt.Println("Salvando usuário", u.nome, "com idade", u.idade) // método salvar
}

func main() {
	usuario1 := usuario{"Lucas", 25}
	fmt.Println(usuario1) // imprime o usuário
	usuario1.salvar()     // chama o método salvar do usuário 1
}
