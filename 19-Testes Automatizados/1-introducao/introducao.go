package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("rua das flores")
	fmt.Printf("O tipo de endereço é: %s\n", tipoEndereco)

}
