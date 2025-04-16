package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)      // Cria um canal de string
	go escrever("Olá Mundo", canal) // Goroutine

	fmt.Println("Depois da funçao escrever começar a ser executada") // Imprime mensagem antes de receber do canal

	for {
		mensagem := <-canal   // Recebe a mensagem do canal
		fmt.Println(mensagem) // Imprime a mensagem recebida
	}

}

func escrever(texto string, canal chan string) {
	time.Sleep(5 * time.Second) // Espera 5 segundos antes de enviar a mensagem
	for i := 0; i < 5; i++ {
		canal <- texto // Envia a mensagem para o canal
		time.Sleep(time.Second)
	}
}
