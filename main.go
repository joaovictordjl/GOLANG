package main

import "fmt"

func main ()  {
	fmt.Println("Hello World, my name is João Victor. This is my first code in the lenguage of programming Golang")
}  

//PRINTLN IMPRIME E PULA UMA LINHA

//Package main- executará a função main antes de tudo, ao dar run no programa

/*
	PACOTES:
	
	São conjuntos de funções. Um pacote, tem várias funções.
		-> fmt é um pacote e usa a função Println
	
	Para importar um pacote é simples:
		-> import "nomeDoPacote"
			Caso tenha mais de um pacote para importar, faça assim:
				-> import (
					"pacote1"
					"pacote2"
					"pacote3"
				)
*/