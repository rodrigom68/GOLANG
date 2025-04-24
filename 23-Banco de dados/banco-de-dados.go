package main

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

func main() {
	// Conecta ao banco de dados PostgreSQL
	connStr := "host=127.0.0.1 port=5432 user=testeGo password=123Go dbname=fieldist sslmode=disable"	// String de conexão com o banco de dados
	db, err := sql.Open("postgres", connStr)	// Abre a conexão com o banco de dados	
	if err != nil {
		log.Fatal(err)	// Se houve erro, imprime o erro
	}	// Se não houve erro, imprime a mensagem de sucesso
	defer db.Close()	// Fecha a conexão com o banco de dados quando a função main terminar