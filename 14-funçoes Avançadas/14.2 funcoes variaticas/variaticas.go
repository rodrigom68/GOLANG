package main

import "fmt"

func soma(numeros ...int) {
	fmt.Println("Números:", numeros)
}

func main() {
	soma(1, 2, 3, 4, 5)
}
