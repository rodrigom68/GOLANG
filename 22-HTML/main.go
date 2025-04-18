package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string // Nome do usuário
	Email string // Email do usuário
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

	templates = template.Must(template.ParseGlob("*.html")) // Carrega os templates HTML.

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{ // Cria um usuário com nome e email.
			Nome:  "Rafael",                 // Nome do usuário
			Email: "rafa21534548@gmail.com", // Email do usuário
		} // Cria um usuário com nome e email.

		templates.ExecuteTemplate(w, "home.html", u) // Executa o template home.html e escreve na resposta.
	}) // Função anônima que recebe um ResponseWriter e um Request como parâmetros.

	fmt.Println("Servidor rodando na porta 5000...") // Imprime mensagem no console.

	log.Fatal(http.ListenAndServe(":5000", nil)) // Inicia o servidor HTTP na porta 8080 e aguarda requisições.

}
