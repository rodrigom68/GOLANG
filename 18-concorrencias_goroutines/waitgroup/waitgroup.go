package main

import (
	"fmt"
	"time"
)

func main() {
	escrever("Olá Mundo") // Goroutine
	escrever("Programando em Go")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
