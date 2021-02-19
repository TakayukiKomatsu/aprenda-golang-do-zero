package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int16 = 100
	fmt.Println(numero)

	var numero2 int64 = 100000
	fmt.Println(numero2)

	var numero3 uint32 = 100
	fmt.Println(numero3)

	//ALIAS
	//Int32 = rune
	var numero4 rune = 123456
	fmt.Println(numero4)

	//uint0 = btye
	var numero5 byte = 123
	fmt.Println(numero5)

	//Numeros reais

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64
	fmt.Println("numeroReal2", numeroReal2)

	numeroReal3 := 1234.123
	fmt.Println(numeroReal3)

	//Strings
	var str string = "qualquerTexto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//Char

	char := 'A'
	fmt.Println(char) //Retrona o numero na tabela ascii

	//valor Zero == valor incial caso não seja declarado o valor

	var texto int
	fmt.Println(texto)

	//Booleano

	var boolean bool = true //O valor zero é false
	fmt.Println(boolean)

	//Erro => é um tipo de dados

	var err error = errors.New("Erro ")

	fmt.Println(err)

}

//Tipos básicos no go - há dois tipos numericos, sendo int  e float

/*
	Numeros Inteiros:
		- int8
		- int16
		- int32
		- int64


		int = o tamanho da variavel depende da arquitetura do computador, se for 64 bits, será um int64,
		e no caso de inferencia de tipos , é utilizada int
	Unsigned int:
		respeita o mesmo tamanho dos int, porem, sem sinal


		alias:
		int32 = rune
		uint8 = byte
*/

/*
	Numeros Reais:
	- float32
	- float64
*/
