package main

import (
	"fmt"
	"time"
)

func main() {

	canal := escrever("Olá Mundo") // Goroutine

	for i := 0; i < 5; i++ { // Loop para receber mensagens do canal
		fmt.Println(<-canal) // Imprime a mensagem recebida
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string) // Cria um canal de string

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto) // Envia a mensagem para o canal
			time.Sleep(time.Millisecond * 500)                // Espera milisegundo antes de enviar a próxima mensagem
		}
	}()

	return canal // Retorna o canal

}
