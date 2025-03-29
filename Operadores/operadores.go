package main

import "fmt"

func main() {
	//ARITMÉTICOS
	soma := 10 + 5
	subtracao := 20 - 10
	divisao := 20 / 10
	multiplicacao := 2 * 5
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	fmt.Println("")

	/*
		não é possivel realizar operações com variáveis que tenham o TIPO diferente. Exemplo abaixo:
			var numero1 int16 = 10
			var numero2 int8 = 20
			somando := numero1 + numero2
			 println(somando)
	*/

	//OPERADORES DE ATRIBUÍÇÃO

	var variavel1 string = "variavel1"
	variavel2 := "variavel2"
	fmt.Println(variavel1, variavel2)
	fmt.Println("")


	//OPERADORES RELACIONAIS  
	fmt.Println(10 > 5) //verifica se é maior
	fmt.Println(10 <= 6) //verifica se é menor ou igual
	fmt.Println(5 == 8) // verifica se é igual
	fmt.Println(2 <= 4) // verifica se é menor ou igual 
	fmt.Println(2 < 3) // verifica se é menor
	fmt.Println(15 != 10) // verifica se é diferente
	fmt.Println("")

	//OPERADORES LÓGICOS

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //TODOS PRECISAM SER VERDADEIROS PARA SER TRUE
	fmt.Println(verdadeiro || falso) //APENAS UM PRECISA SER VERDADEIRO PARA SER TRUE
	fmt.Println(!verdadeiro, !falso) //OPERADOR DE NEGAÇÃO, INVERTE O VALOR
	fmt.Println("") 

	//OPERADORES UNÁRIOS
	numero := 1
	numero++ //incrementa uma unidade
	numero-- //decrementa 
	numero += 10 // numero = numero + 10
	fmt.Println("numero é ", numero) //exemplo de concatenação
}