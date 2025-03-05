package main

import (
	"fmt"             // pacote de formatação
	"modulo/auxiliar" // importando o pacote auxiliar
)

func main() {
	auxiliar.Escrever()                       // chamando a função do pacote auxiliar
	fmt.Println("Escrevendo no arquivo main") // imprime o texto na tela
}
