package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}
func divisao(a float64, b float64) float64 {
	if b == 0 {
		panic("O denominador não pode ser 0")

	}
	return a / b
}
func multiplicacao(a int, b int) int {
	return a * b
}

func main() {
	// Operadores aritméticos
	// + - * / %
	// ++ --
	// += -= *= /= %=
	// == != > < >= <=
	// && || !
	fmt.Println(soma(3, 4))
	fmt.Println(divisao(3, 4))
	fmt.Println(multiplicacao(3, 4))

}
