package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`  // Nome do cachorro
	Idade int    `json:"idade"` // Idade do cachorro
	Raca  string `json:"raca"`  // Raça do cachorro
}

func main() {

	cachorroEmJSON := `{"nome":"Rex","idade":5,"raca":"Labrador"}` // Cria um cachorro com nome Rex, idade 5 e raça Labrador
	var c cachorro                                                 // Cria um cachorro vazio

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro) // Imprime o erro
	} // Se não houve erro, imprime o cachorro em JSON

	fmt.Println(c) // Imprime o cachorro em JSON

	// Converte o JSON para cachorro

	cachorro2 := `{"idade":"5","nome":"Caramelo","raca":"Pitbull"}`

	cachorroEmJSON2 := make(map[string]string) // Cria um mapa com os dados do cachorro

	if erro := json.Unmarshal([]byte(cachorro2), &cachorroEmJSON2); erro != nil {
		log.Fatal(erro) // Imprime o erro
	} // Se não houve erro, imprime o cachorro em JSON

	fmt.Println(cachorroEmJSON2) // Imprime o cachorro em JSON como bytes

	// Converte o JSON para cachorro
	// Verifica se houve erro na conversão

}
