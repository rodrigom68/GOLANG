package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // Cria um WaitGroup

	waitGroup.Add(2) // Adiciona duas goroutines ao WaitGroup

	go func() {
		escrever("Olá Mundo") // Goroutine
		waitGroup.Done()      // Indica que a goroutine terminou
	}()
	go func() {
		escrever("Programando em Go") // Goroutine
		waitGroup.Done()              // Indica que a goroutine terminou
	}()

	waitGroup.Wait() // Aguarda todas as goroutines terminarem

	escrever("Olá Mundo") // Goroutine
	escrever("Programando em Go")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
