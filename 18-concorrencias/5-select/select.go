package main

import "time"

func main() {
	canal1, canal2 := make(chan string), make(chan string) // Cria dois canais de string

	go func() {

		for {
			time.Sleep(500 * time.Millisecond) // Espera 2 segundos antes de enviar a mensagem
			canal1 <- "Canal 1"                // Envia a mensagem para o canal 1

		}

	}()

	go func() {

		for {
			time.Sleep(2 * time.Second) // Espera 2 segundos antes de enviar a mensagem
			canal2 <- "Canal 1"         // Envia a mensagem para o canal 1

		}

	}()

	for {

		select { // Seleciona entre os canais
		case mensagemCanal1 := <-canal1: // Recebe a mensagem do canal 1
			println(mensagemCanal1) // Imprime a mensagem recebida do canal 1
		case mensagemCanal2 := <-canal2: // Recebe a mensagem do canal 2
			println(mensagemCanal2) // Imprime a mensagem recebida do canal 1
	// 	mensagemCanal1 := <-canal1 // Recebe a mensagem do canal 1
	// 	println(mensagemCanal1)    // Imprime a mensagem recebida do canal 1

	// 	mensagemCanal2 := <-canal2 // Recebe a mensagem do canal 2
	// 	println(mensagemCanal2)    // Imprime a mensagem recebida do canal 1
	// }

}
