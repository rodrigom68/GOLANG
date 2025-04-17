package main

import (
	"fmt"
	"time"
)

func main() {

	canal := multiplexar(escrever("Olá Mundo"), escrever("Programando em Go")) // Cria um canal multiplexado
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal) // Recebe a mensagem do canal multiplexado

	}

}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDesaida := make(chan string)
	go func() {
		for {
			select { // Seleciona o canal que está pronto para enviar ou receber
			case mensagem := <-canalDeEntrada1:
				canalDesaida <- mensagem // Envia a mensagem do canal1 para o canal multiplexado
			case mensagem := <-canalDeEntrada2:
				canalMultcanalDesaidaiplexado <- mensagem // Envia a mensagem do canal2 para o canal multiplexado
			}
		}
	}()

	return canalDesaida // Retorna o canal multiplexado
}

func escrever(texto string) <-chan string {
	canal := make(chan string) // Cria um canal de string

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto) // Envia a mensagem para o canal
			time.Sleep(time.Millisecond * 500)                // Espera milisegundo antes de enviar a próxima mensagem
		}
	}()

	return canal
} // Retorna o canal
