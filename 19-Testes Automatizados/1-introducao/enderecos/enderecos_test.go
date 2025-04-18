// TESTE DE UNIDADE

package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	tipoEsperado     string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua das Flores", "Rua"},
		{"Avenida das Flores", "Avenida"},
		{"Praça das Flores", "Praça"},
		{"Rodovia das Flores", "Rodovia"},
		{"Estrada das Flores", "Estrada"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.tipoEsperado {
			t.Errorf("Esperado: %s, Recebido: %s", cenario.tipoEsperado, tipoDeEnderecoRecebido)
		}
	}

	// Testa se o tipo de endereço é válido
	enderecoParaTeste := "das flores"
	tipoEsperado := "Rua"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoEsperado {
		t.Errorf("Esperado: %s, Recebido: %s", tipoEsperado, tipoDeEnderecoRecebido)
	}
}
