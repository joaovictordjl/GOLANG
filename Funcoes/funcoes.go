package main

import "fmt"

//função que realiza soma de 2 valores, que são do tipo int8
func Somar(n1, n2 int8) int8 {
	return n1 + n2
}

//função que realiza multiplicação de 2 valores
func Multiplicar(n1, n2 int8) int8 {
	return n1 * n2
}

func main() {
	soma := Somar (5, 2)
	fmt.Println(soma)
	mutiplicacao := Multiplicar(5,2)
	fmt.Println(mutiplicacao)
	
}
