package main

import "fmt"

func main() {
	// A função init é executada antes da função main
	fmt.Println("Executando a função init")
	// A função main é executada depois da função init
	fmt.Println("Executando a função main")
}
