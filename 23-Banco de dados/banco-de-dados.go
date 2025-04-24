package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
	// _ "github.com/go-sql-driver/mysql" // Importa o driver do MySQL
)

func main() {
	// Conecta ao banco de dados PostgreSQL
	connStr := "host=127.0.0.1 port=5432 user=testeGo password=123Go dbname=fieldist sslmode=disable" // String de conexão com o banco de dados
	db, err := sql.Open("postgres", connStr)                                                          // Abre a conexão com o banco de dados
	if err != nil {
		log.Fatal(err) // Se houve erro, imprime o erro
	} // Se não houve erro, imprime a mensagem de sucesso
	defer db.Close() // Fecha a conexão com o banco de dados quando a função main terminar

	//testa a conexão com o banco de dados
	err = db.Ping() // Testa a conexão com o banco de dados
	if err != nil {
		log.Fatal(err) // Se houve erro, imprime o erro
	} // Se não houve erro, imprime a mensagem de sucesso
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!", db) // Imprime a mensagem de sucesso
}
