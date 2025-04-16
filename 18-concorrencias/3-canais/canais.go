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
		mensagem, aberto := <-canal // Recebe a mensagem do canal
		if !aberto {                // Verifica se o canal está aberto
			break
		}
		fmt.Println(mensagem) // Imprime a mensagem recebida
	}

	fmt.Println("Canal fechado") // Imprime mensagem após o canal ser fechado

}

func escrever(texto string, canal chan string) {
	time.Sleep(5 * time.Second) // Espera 5 segundos antes de enviar a mensagem
	for i := 0; i < 5; i++ {
		canal <- texto // Envia a mensagem para o canal
		time.Sleep(time.Second)
	}
	close(canal) // Fecha o canal após enviar todas as mensagens
}
