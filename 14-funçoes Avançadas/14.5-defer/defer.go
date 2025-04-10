package main

import "fmt"

func funcao1() {
	fmt.Println("Executando funcao 1")
}

func funcao2() {
	fmt.Println("Executando funcao 2")
}

func main() {
	fmt.Println("Iniciando o programa")
	// Adiando funçao1
	defer funcao1() //defer = adiar execução da função
	funcao2()

	fmt.Println("Finalizando o programa")
}
