package app

import "github.com/urfave/cli"

// Gerar vai retornar uma aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()                                     // cria uma nova aplicação de linha de comando
	app.Name = "Aplicação de linha de comando"              // nome da aplicação
	app.Usage = "Busca IPS e Nomes de Servidor na internet" // descrição da aplicação

	app.Commands = []cli.Command{ // comandos da aplicação
		{
			Name:  "ip", // nome do comando
			Usage: "Busca o IP de um servidor na internet", // descrição do comando
			Flags: []cli.Flag{ // flags do comando
				cli.StringFlag{
					Name:  "servidor", // nome da flag
					Value: "frigo-data.com.br", // valor padrão da flag				  
				},
			},
			Action: buscarIps, // ação do comando
		},
	}

	return app // retorna a aplicação de linha de comando
}

func buscarIps(c *cli.Context) { 	
	host := c.String("host") // pega o valor da flag servidor

	ips, erro := net.LookupIP(host) // busca o IP do servidor na internet
	if erro != nil { // se houver erro, imprime o erro
		log.Fatal(erro) // registra o erro e encerra a aplicação
}
