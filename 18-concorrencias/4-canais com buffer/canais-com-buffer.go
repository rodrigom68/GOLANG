package main

func main() {
	canal := make(chan string, 2) // Cria um canal de string com buffer de 2

	canal <- "Olá Mundo" // Envia a mensagem para o canal

	mensagem := <-canal // Recebe a mensagem do canal
	println(mensagem)   // Imprime a mensagem recebida

}
