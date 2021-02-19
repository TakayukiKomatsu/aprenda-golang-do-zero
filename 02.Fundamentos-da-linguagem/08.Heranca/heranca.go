package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint16
}

type estudante struct {
	pessoa //Está "pegando" os valores de pessoa e atribuindo a struct estudante, como se fosse copiar e colar

	curso     string
	faculdade string
}

func main() {

	fmt.Println("Heranca")

	p1 := pessoa{"João", "Pedro", 20, 178}

	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
}
