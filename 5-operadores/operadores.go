package main	

func main()  {
	// Operadores aritméticos
	// + - * / %
	// ++ --
	// += -= *= /= %=
	// == != > < >= <=
	// && || !

	func soma(a int, b int) int {
		return a + b
	}
	func divisao(a int, int b) int {
		if b == 0 {	
			panic("O denominador não pode ser 0")
			
		}
		return a / b
	}
	func multiplicacao(a int, b int) int {
		return a * b
	}
 }