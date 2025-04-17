package main

func main() {


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