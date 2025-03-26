package main

import "fmt"

func main() {

	primeiroNome := "João "
	fmt.Println(primeiroNome)


	//DECLARAR DE FORMA EXPLÍCITA: var  +nomeDaVariavel + tipoDaVariável = valor

	//FORMA IMPLÍCITA(INFERENCIA DE TIPO): nomeDaVariavel := valor


	//DECLARAR MAIS DE UMA POR VEZ(FORMA EXPLÍCITA):

		var (
		variavel1 string = "lalala"
		variavel2 string = "lelele"
		)

		fmt.Println(variavel1, variavel2)


	//DECLARAR MAIS DE UMA(DA FORMA DE INFERENCIA DE TIPOS)
		
		variavel3, variavel4 := "lilili", "lololo"

		fmt.Println(variavel3, variavel4)

	//CONSTANTES:

		const fraseImportante string = "Eu amo a Gabrielle"
		
		fmt.Println(fraseImportante)

}