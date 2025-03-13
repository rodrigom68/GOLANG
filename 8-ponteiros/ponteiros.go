package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)
	variavel2++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória

	var variavel3 int = 100
	var ponteiro *int

	variavel3 = 150
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // desreferenciando o ponteiro

}
