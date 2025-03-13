package main

import "fmt"

func main() {
	type Pessoa struct {
		nome  string
		idade int
	}

	pessoa := Pessoa{"João", 20}
	fmt.Println(pessoa)
}
