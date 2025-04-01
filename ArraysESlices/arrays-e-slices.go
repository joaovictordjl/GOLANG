package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

					//ARRAY pouco utilizado em GO

	//ARRAY EM GOLANG É MUITO INFLEXÍVEL
		//NÃO PODE SER ADICIONADO VALORES POSTERIORMENTE AO ARRAY
			
			array2 := [5]string{"POSIÇÃO 1", "POSIÇÃO 2", "POSIÇÃO 3", "POSIÇÃO 4", "POSIÇÃO 5"}
			fmt.Println(array2)

		//selecionar um dado expecífico de um array:
			fmt.Println(array2[3])

		/*
			A ordem das posições dos dados nos arrays, iniciam contando a partir do 0. Então 0 seria a primeira posição. 
				Como selecionei 1, que seria a 2ª posição, ele me retorna "POSIÇÃO 2".
		*/

				//SLICES- BASTANTE UTILIZADO EM GO(É UMA "FATIA" DO ARRAY)
		
		//Não precisa colocar o tamanho(quantidade de dados) nos colchetes 
			slice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			fmt.Println(slice)

		//adicionar um dado atualizando ao mesmo tempo

			slice = append(slice, 11)
			fmt.Println(slice)

		//REFERENCIAR O ARRAY através do slice

			//Aqui eu busco dados num intervalo entre 1º e 4º posição
				slice2 := array2[1:4]
				fmt.Println(slice2)

}
