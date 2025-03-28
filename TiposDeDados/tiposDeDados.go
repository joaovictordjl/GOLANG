package main

import "fmt"

func main() {

	//INT

		//int8, int16, int32, int64

		//FORMA EXPLÍCITA

			var numero int64= 10000000000000
			fmt.Println(numero)

		/*

		IMPLÍCITA
			A forma implícita, irá se basear nos bits do sistema do seu computador. Se ele for 64bits, o int será executado como int64 e assim sucessivamente.

			EXEMPLO
			 |
			 |
			 V
		*/
			numero2 := 10000
			fmt.Println(numero2)

	//UINT-> NÃO PODE USAR SINAIS

			//AQUI DARIA UM ERRO var numero3 uint = -5

			var numero3 uint = 5 //AQUI NÂO DARIA ERRO
			fmt.Println(numero3)

	//NUMEROS DECIMAIS

		//EXPLÍCITA
			var numeroDecimal float32 = 255.5
			fmt.Println(numeroDecimal)

			var numeroDecimal2 float64 = 12345678.08
			fmt.Println(numeroDecimal2)

		//IMPLÍCITA baseia-se nos bits do meu sistema
			numeroDecimal3 := 25.4
			fmt.Println(numeroDecimal3)


	//STRING = SEMPRE ESTAR ENTRE ""
		
		texto1 := "Texto aqui é uma string" 
		fmt.Println(texto1)


	//BOOLEANO

		var booleano bool = true
		fmt.Println(booleano)

	//ERRO TAMBÉM É UM TIPO

		var erro error
		fmt.Println(erro)
}