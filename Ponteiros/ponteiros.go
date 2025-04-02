package main

import "fmt"

func main() {
	fmt.Println("Aprendendo ponteiros")

		//É UMA REFERENCIA DE MEMÓRIA- SERVE PARA MODIFICAR, COMPARTILHAR OU ACESSAR O VALOR EM DIFERENTES PARTES DO PROGRAMA
	
	var variavel3 int 
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 //usar o &, indica o endereço de memória da variável(onde ela foi armazenada)
	fmt.Println(variavel3, ponteiro)

	/*
		Neste caso, ponteiro não está recebendo o valor 100 da variável3, ele apenas diz o endereço de memória dela(onde o valor 100 está armazenado)
	*/


	//para ler o valor encontrado no endereço de memória da variável, faça assim:
		fmt.Println(variavel3, *ponteiro)


		//mais um exemplo : Estante de gavetas

			idade := 10       // Gaveta com etiqueta "idade" guarda o valor 10
			ponteiroIdade := &idade // "ponteiroIdade" guarda o número da gaveta(endereço/local) da variável "idade"

			fmt.Println(idade)        // Imprime 10
			fmt.Println(ponteiroIdade) // Imprime o número da gaveta (endereço/local)

			*ponteiroIdade = 20     // Muda o valor que está dentro da gaveta para 20

			fmt.Println(idade)        // Imprime 20



			numeroDaSorte := 21
			ondeEstaONumeroDaSorte := &numeroDaSorte

			fmt.Println(numeroDaSorte, ondeEstaONumeroDaSorte)

			*ondeEstaONumeroDaSorte = 22  //AQUI EU MODIFIQUEI MEU NÚMERO DA SORTE, ATRAVÉS DO PONTEIRO
			fmt.Println(numeroDaSorte)

}

