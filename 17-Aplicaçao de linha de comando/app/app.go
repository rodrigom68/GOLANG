package app

import "github.com/urfave/cli"

// Gerar vai retornar uma aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()                                     // cria uma nova aplicação de linha de comando
	app.Name = "Aplicação de linha de comando"              // nome da aplicação
	app.Usage = "Busca IPS e Nomes de Servidor na internet" // descrição da aplicação

	return app // retorna a aplicação de linha de comando

}
