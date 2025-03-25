package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 0

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero < 5 {
		fmt.Println("Menor que 5")
	} else {
		fmt.Println("Entre 5 e 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	}
}
