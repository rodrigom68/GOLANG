package main

import (
	"fmt"
	"time"
)

func main() {
	// Estruturas de repetição (loops)
	// for, while, do while
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		return
	// 	}
	// 	fmt.Println(i)
	// }
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("incrementando i =", i)
	// 	time.Sleep(time.Second)

	// }
	// fmt.Println("i =", i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("incrementando j =", j)
	// 	time.
	// 		Sleep(time.Second)
	// }

	// nomes := [3]string{"João", "Maria", "José"}

	// for indice, nomes := range nomes {
	// 	fmt.Println("índice:", indice, "nome:", nomes)
	// 	time.Sleep(time.Second)
	// }

	// for _, nome := range nomes {
	// 	fmt.Println("nome:", nome)
	// 	time.Sleep(time.Second)
	// }

	// for indice, letra := range "Programação" {
	// 	fmt.Println("índice:", indice, string(letra))
	// 	time.Sleep(time.Second)
	// }

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
		"idade":     "30",
	}

	for chave, valor := range usuario {
		fmt.Println("chave:", chave, valor)
		time.Sleep(time.Second)
	}
}
