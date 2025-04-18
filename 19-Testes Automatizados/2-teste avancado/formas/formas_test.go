package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{base: 10, altura: 15} // criando um retângulo

		areaEsperada := float64(150) // área esperada do retângulo
		areaRecebida := ret.Area()   // chamando o método area do retângulo

		if areaEsperada != areaRecebida { // verificando se a área recebida é igual à área esperada
			t.Fatalf("Área do retângulo: esperada %v, recebida %v", areaEsperada, areaRecebida) // erro na verificação
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{raio: 10} // criando um círculo

		areaEsperada := float64(math.Pi * math.Pow(10, 2)) // área esperada do círculo
		areaRecebida := circ.Area()                        // chamando o método area do círculo

		if areaEsperada != areaRecebida { // verificando se a área recebida é igual à área esperada
			t.Fatalf("Área do círculo: esperada %v, recebida %v", areaEsperada, areaRecebida) // erro na verificação
		}
	})

}
