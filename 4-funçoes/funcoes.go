package main

func somar(a int, b int) int {
	return a + b
}

func calculosMatematicos(a int, b int) (int, int) {
	soma := a + b
	subtracao := a - b
	return soma, subtracao
}

func main() {
	resultado := somar(3, 4)
	println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(3, 4)
	println(resultadoSoma, resultadoSubtracao)
}
