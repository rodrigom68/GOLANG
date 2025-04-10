package main

import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool {
	notaFinal := (n1 + n2) / 2
	// notaFinal := (n1 + n2) / 2.0 // ou float64(n1+n2)/2.0
	if notaFinal > 6 {
		return true
	} else if notaFinal < 6 {
		return false
	}

	panic("Nota final é exatamente 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6)) // true
	fmt.Println("Pos Execução")
}
