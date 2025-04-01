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
	i := 0
	for i < 10 {
		i++
		fmt.Println("incrementando i =", i)
		time.Sleep(time.Second)

	}
	fmt.Println("i =", i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("incrementando j =", j)
		time.
			Sleep(time.Second)
	}

	nomes := [3]string{"João", "Maria", "José"}

	for indice, nomes := range nomes {
		fmt.Println("índice:", indice, "nome:", nomes)
		time.Sleep(time.Second)
	}
}
