package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome  string
	idade uint8

	endereco endereco
}

func main() {

	fmt.Println("Arquivo Structs")

	var u usuario

	u.nome = "randon"
	u.idade = 40

	fmt.Println(u)

	enderecoExemplo := endereco{"Numero dos bobos", 0}

	fmt.Println("-------------------")

	user2 := usuario{"randon", 21, enderecoExemplo}

	fmt.Println(user2)

	fmt.Println("-------------------")

	user3 := usuario{idade: 21}
	fmt.Println(user3)

}
