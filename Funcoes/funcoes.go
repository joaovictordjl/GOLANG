package main

import "fmt"


//função que realiza calculos matemáticos com dois valores
func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
	
}
/*
	O PRIMEIRO parêntese são os parâmetros da entrada da função(n1 e n2)
		O SEGUNDO parêntese significa os tipos dos retornos da função, que no nosso caso são valores do tipo int8, respectivamente: soma, subtracao e multiplicação.
*/

func main() {
	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao:= calculosMatematicos(10, 2)
	fmt.Println(resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao)
}
