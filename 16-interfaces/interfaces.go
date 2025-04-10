package main

import (
	"fmt"
	"math"
)

type forma interface { // interface forma
	area() float64 // método area que retorna o valor da área da forma
}

type retangulo struct {
	base   float64 // base do retângulo
	altura float64 // altura do retângulo
}

func (r retangulo) area() float64 { // método area do retângulo
	return r.base * r.altura // retorna a área do retângulo
}

func (c circulo) area() float64 { // método area do círculo
	return math.Pi * math.Pow(2, c.raio) // retorna a área do círculo
}

type circulo struct {
	raio float64 // raio do círculo
}

func main() {
	ret := retangulo{base: 10, altura: 15} // criando um retângulo
	cir := circulo{raio: 10}               // criando um círculo

	fmt.Println("Área do retângulo:", ret.area()) // chamando o método area do retângulo
	fmt.Println("Área do círculo:", cir.area())   // chamando o método area do círculo

	var f forma                             // criando uma variável do tipo forma
	f = ret                                 // atribuindo o retângulo à variável forma
	fmt.Println("Área da forma:", f.area()) // chamando o método area da forma

	f = cir                                 // atribuindo o círculo à variável forma
	fmt.Println("Área da forma:", f.area()) // chamando o método area da forma
}
