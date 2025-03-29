package main

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main, textando o mod")
	fmt.Println("teste")

	auxiliar.Escrever()


	//pacote externo para validação de email, explicação para importar estará abaixo(PE)
	checkmail.ValidateFormat("victorjhon779@gmail.com")
}

//Pacotes são conjuntos de funções ou arquivos de uma pasta

/* 
Para usar mais de um pacote em um projeto, utiliza-se o módulo
	Para criar um módulo:
		go mod init modulo
	Modulo é onde os pacotes ficam armazenados


Após criar um módulo e escrever algum código, escreva no terminal:
	go build
		Isso irá gerar um arquivo binário, que compila todos os arquivos do seu projeto em um só, que seria esse binário.
			É possivel executar seu projeto rodando apenas esse arquivo binário, da seguinte forma:
				./modulo


Cada pacote precisa estar em uma pasta diferente, ou seja, sempre que for criar um pacote é necessário criar uma pasta para ele.


(PE) Importar Pacote Externo pelo terminal:

	go get urlDoPacote 
		
		Após importar o pacote, ele deve ser adicionado na aba de import ()
		Posteriormente, ele precisa ser chamado na função, ex:
			checkmail.ValidateFormat("") que foi o caso que utilizamos

*/