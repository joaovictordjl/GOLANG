package main

import "fmt"

//AGRUPA VARIAVEIS DE DIFERENTES TIPOS EM UM ÚNICO NOME
//neste caso, usuário seria um "TIPO" onde essas variáveis seriam armazenadas
type usuario struct {
	nome string
	idade uint
}

func main() {
	fmt.Println("Brincando com structs")

	//CRIAÇÃO DA VARIÁVEL COM O STRUCT CRIADO.
		//o nome da variável é usuário1 e ela é do tipo usuario(struct)
			usuario1 := usuario {"João", 22}
			fmt.Println(usuario1)
	//os dois valores que ela recebe, são respectivos dos valores criados no struct(nome e idade)

		/*
			QUANDO USAR: Quando precisar de mais de um tipo de dado para uma unica variável.
			Como exeplificamos, precisavamos de um usuário e que armazenasse a idade e o nome dele
		*/
}