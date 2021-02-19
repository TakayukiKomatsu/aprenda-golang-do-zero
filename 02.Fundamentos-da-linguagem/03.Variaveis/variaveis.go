package main

import "fmt"

func main() {

	/*há duas formas de declarar as variaveis, de forma implícita e explicita*/
	var variavel1 string = "Variavel 1" /*Declaração de forma explicita*/
	fmt.Println(variavel1)

	variavel2 := "Variavel 2" /*Declaração de forma implicita - Inferencia de tipos*/
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	/*Troca de valores*/

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
