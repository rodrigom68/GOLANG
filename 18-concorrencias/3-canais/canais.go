package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)      // Cria um canal de string
	go escrever("Olá Mundo", canal) // Goroutine

	mensagem := <-canal   // Recebe a mensagem do canal
	fmt.Println(mensagem) // Imprime a mensagem recebida
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Envia a mensagem para o canal
		time.Sleep(time.Second)
	}
}
