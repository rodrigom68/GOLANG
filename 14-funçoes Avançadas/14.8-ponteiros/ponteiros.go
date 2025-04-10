package main

import "fmt"

func inverterSinal(numero int) int {
	if numero > 0 {
		return -numero
	}
	return -numero
}

func main() {
	numero := 20
	numeInvertido := inverterSinal(numero)
	fmt.Println("Número original:", numero)
	fmt.Println("Número invertido:", numeInvertido)
}
