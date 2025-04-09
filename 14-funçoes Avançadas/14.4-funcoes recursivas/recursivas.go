package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funcoes Recursivas")

	// 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946

	posicao := uint(10)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println("Fibonacci de", i, "é", fibonacci(i))

		fmt.Println("Fibonacci de", posicao, "é", fibonacci(posicao))
	}
}
