package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar uma aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()                                     // cria uma nova aplicação de linha de comando
	app.Name = "Aplicação de linha de comando"              // nome da aplicação
	app.Usage = "Busca IPS e Nomes de Servidor na internet" // descrição da aplicação

	flags := []cli.Flag{ // flags do comando
		cli.StringFlag{
			Name:  "host",              // nome da flag
			Value: "frigo-data.com.br", // valor padrão da flag
		},
	}

	app.Commands = []cli.Command{ // comandos da aplicação
		{
			Name:   "ip", // nome do comando
			Usage:  "Busca o IP de um servidor na internet",
			Flags:  flags,     // descrição do comando
			Action: buscarIps, // ação do comando
		},
		// comando para buscar o nome do servidor na internet
		{
			Name:   "servidores", // nome do comando
			Usage:  "Busca o nome de um servidores na internet",
			Flags:  flags,            // descrição do comando
			Action: buscarServidores, // ação do comando
		},
	}

	return app // retorna a aplicação criada
}

func buscarServidores(c *cli.Context) {
	host := c.String("host") // pega o valor da flag servidor

	servidores, erro := net.LookupNS(host) // busca o nome do servidor na internet
	if erro != nil {                       // se houver erro, imprime o erro
		log.Fatal(erro) // registra o erro e encerra a aplicação
	}

	for _, servidor := range servidores { // para cada servidor encontrado
		fmt.Println(servidor.Host) // imprime o servidor
	}
}

func buscarIps(c *cli.Context) {
	host := c.String("host") // pega o valor da flag servidor

	ips, erro := net.LookupIP(host) // busca o IP do servidor na internet
	if erro != nil {                // se houver erro, imprime o erro
		log.Fatal(erro) // registra o erro e encerra a aplicação
	}

	for _, ip := range ips { // para cada IP encontrado
		fmt.Println(ip) // imprime o IP
	}
}
