package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) { // Cria uma rota para a URL /home
	w.Write([]byte("Hello World!")) // Escreve "Hello World!" na resposta.
} // Função anônima que recebe um ResponseWriter e um Request como parâmetros.

func usuario(w http.ResponseWriter, r *http.Request) { // Cria uma rota para a URL /home
	w.Write([]byte("Carrega pagina de usuarios!")) // Escreve "Hello World!" na resposta.
}

func main() {

	// HTTP é um protocolo de comunicação entre cliente e servidor.
	// Ele é baseado em requisições e respostas, onde o cliente envia uma requisição e o servidor responde com os dados solicitados.

	// Request  - Response - Status Code - Headers - Body

	// Rotas - Endpoints - URL - URI - Query String - Path - Method - GET - POST - PUT - DELETE - PATCH
	// URL - Uniform Resource Locator - Endereço de um recurso na web.
	// URI - Uniform Resource Identifier - Identificador de um recurso na web.
	// method - Método HTTP utilizado na requisição.
	// GET - Método utilizado para obter dados de um recurso.
	// POST - Método utilizado para enviar dados para um recurso.
	// PUT - Método utilizado para atualizar dados de um recurso.
	// DELETE - Método utilizado para remover dados de um recurso.

	http.HandleFunc("/home", home) // Escreve "Hello World!" na resposta.

	http.HandleFunc("/usuario", usuario) // Escreve "Carrega pagina de usuarios!" na resposta.

	log.Fatal(http.ListenAndServe(":5000", nil)) // Inicia o servidor HTTP na porta 8080 e aguarda requisições.

}
