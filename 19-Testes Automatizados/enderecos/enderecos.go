package enderecos

import "strings"

func tipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true

		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}
	return "tipo inválido"

}
