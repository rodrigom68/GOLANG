package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println("Números:", numeros)
}

func main() {
	soma(1, 2, 3, 4, 5)
	soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
