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
}
