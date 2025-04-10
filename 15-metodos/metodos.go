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

func (u usuario) maiorDeIdade() bool { // método maiorDeIdade
	return u.idade >= 18 // método maiorDeIdade
}

func (u *usuario) fazerAniversario() { // método fazerAniversario
	u.idade++                                                                   // método fazerAniversario
	fmt.Println("Fazendo aniversário do usuário", u.nome, "com idade", u.idade) // método fazerAniversario
}

func main() {
	usuario1 := usuario{"Lucas", 25}
	usuario2 := usuario{"Ana", 30}
	usuario3 := usuario{"João", 22}
	fmt.Println(usuario1)                // imprime o usuário
	usuario1.salvar()                    // chama o método salvar do usuário 1
	usuario2.salvar()                    // chama o método salvar do usuário 2
	usuario3.salvar()                    // chama o método salvar do usuário 3
	usuario1.maiorDeIdade()              // chama o método maiorDeIdade do usuário 1
	usuario2.maiorDeIdade()              // chama o método maiorDeIdade do usuário 2
	usuario3.maiorDeIdade()              // chama o método maiorDeIdade do usuário 3
	fmt.Println(usuario1.maiorDeIdade()) // imprime se o usuário 1 é maior de idade

	usuario1.fazerAniversario() // chama o método fazerAniversario do usuário 1
	usuario2.fazerAniversario() // chama o método fazerAniversario do usuário 2
	usuario3.fazerAniversario() // chama o método fazerAniversario do usuário 3
}
