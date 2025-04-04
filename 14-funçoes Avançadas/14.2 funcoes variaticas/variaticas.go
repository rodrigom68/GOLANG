package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total

}

func escrever(texto string, numeros ...int) {
	fmt.Println(texto)
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalDaSoma)

	escrever("Ola Mundo", 1, 2, 3, 4, 5)
}
