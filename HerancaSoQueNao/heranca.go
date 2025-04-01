package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura float32 

}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println(" 'Heranças' em GO")
	
	pessoa1 := pessoa{"Joao", "Victor", 22, 1.78}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Análise e desenvolvimento de sistemas", "Estácio"}
	fmt.Println(estudante1)

		//PARA ACESSAR UM CAMPO EXPECÍFICO DE ALGUM STRUCT, FAÇA ASSIM: 
			fmt.Println(pessoa1.altura)
}