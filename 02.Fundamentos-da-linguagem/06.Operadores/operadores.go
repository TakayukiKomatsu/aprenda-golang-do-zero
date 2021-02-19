package main

import "fmt"

func main() {
	soma := 1 + 2
	fmt.Println(soma)

	substracao := 1 - 2
	fmt.Println(substracao)

	divisao := 10 / 4
	fmt.Println(divisao)

	multiplicacao := 10 * 5
	fmt.Println(multiplicacao)

	restoDaDivisao := 10 % 2
	fmt.Println(restoDaDivisao)

	// No Go, não é possivel fazer operacoes entre tipos diferentes, mesmo sendo numeros.

	//Operadores de Atribuicao

	var variavel1 string = "Hello world"

	variavel2 := "Hello world"

	fmt.Println(variavel1, variavel2)

	//Fim dos operadores de atribuicao

	//Operadores relacionais

	fmt.Println(1 > 2) //Sempre retornam bool

	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS

	fmt.Println(true && false) //and

	fmt.Println(true || false) //or

	fmt.Println(!true)
	//Fim dos operados lógicos

	//OPERADORES UNÁRIOS - não tem prefixado

	numero := 10
	numero++
	numero += 15

	numero--
	numero -= 20
	fmt.Println(numero)

	//OPERADOR TERNÁRIO - não existe

}
