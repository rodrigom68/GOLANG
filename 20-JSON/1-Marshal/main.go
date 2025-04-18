package main

import (
	"bytes"
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
	c := cachorro{
		Nome:  "Rex",      // Nome do cachorro
		Idade: 5,          // Idade do cachorro
		Raca:  "Labrador", // Raça do cachorro
	} // Cria um cachorro com nome Rex, idade 5 e raça Labrador

	// Converte o cachorro para JSON
	cachorroEmJSON, erro := json.Marshal(c) // Converte o cachorro para JSON
	if erro != nil {                        // Verifica se houve erro na conversão
		log.Fatal(erro) // Imprime o erro
	} // Se não houve erro, imprime o cachorro em JSON

	fmt.Println(string(cachorroEmJSON))          // Imprime o cachorro em JSON
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // Imprime o cachorro em JSON como bytes

	c2 := map[string]string{ // Cria um mapa com os dados do cachorro
		"nome":  "Caramelo", // Nome do cachorro
		"idade": "5",        // Idade do cachorro
		"raca":  "Pitbull",  // Raça do cachorro
	} // Cria um cachorro com nome Caramelo, idade 5 e raça Pitbull
	cachorroEmJSON2, erro := json.Marshal(c2) // Converte o cachorro para JSON
	if erro != nil {                          // Verifica se houve erro na conversão
		log.Fatal(erro) // Imprime o erro
	} // Se não houve erro, imprime o cachorro em JSON
	fmt.Println(string(cachorroEmJSON2)) // Imprime o cachorro em JSON

}
