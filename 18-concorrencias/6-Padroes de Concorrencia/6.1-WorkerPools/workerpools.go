package main

func main() {

	tarefas := make(chan int, 45)    // Canal com buffer de 45
	resultados := make(chan int, 45) // Canal com buffer de 45

	go worker(tarefas, resultados) // Inicia a goroutine worker

	for i := 0; i < 45; i++ { // Envia tarefas para o canal
		tarefas <- i
	}
	close(tarefas) // Fecha o canal de tarefas

	for i := 0; i < 45; i++ { // Recebe resultados do canal
		resultado := <-resultados // Recebe o resultado do canal de resultados
		println(resultado)        // Imprime o resultado
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {

	for numero := range tarefas { // Recebe tarefas do canal
		resultados <- fibonacci(numero) // Envia o resultado da tarefa para o canal de resultados
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
