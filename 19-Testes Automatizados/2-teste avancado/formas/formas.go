package formas

import "math"

type Forma interface { // interface forma
	Area() float64 // método area que retorna o valor da área da forma
}

type Retangulo struct {
	base   float64 // base do retângulo
	altura float64 // altura do retângulo
}

func (r Retangulo) Area() float64 { // método area do retângulo
	return r.base * r.altura // retorna a área do retângulo
}

func (c Circulo) Area() float64 { // método area do círculo
	return math.Pi * math.Pow(2, c.raio) // retorna a área do círculo
}

type Circulo struct {
	raio float64 // raio do círculo
}
