package main

import "fmt"

func inverterSinal(numero int) int {
	if numero > 0 {
		return -numero
	}
	return -numero
}

func inverterSinalPonteiro(numero *int) {
	*numero = -(*numero) // Desreferenciação do ponteiro
	// ou numero = -(*numero) // Não funciona, pois altera o valor do ponteiro, não o valor apontado
}

func main() {
	numero := 20
	numeInvertido := inverterSinal(numero)
	fmt.Println("Número original:", numero)
	fmt.Println("Número invertido:", numeInvertido)

	novoNumero := 40
	inverterSinalPonteiro(&novoNumero) // Passando o endereço de memória do número
	fmt.Println("Número original:", novoNumero)
	fmt.Println("Número invertido:", novoNumero) // O número original foi alterado
}
