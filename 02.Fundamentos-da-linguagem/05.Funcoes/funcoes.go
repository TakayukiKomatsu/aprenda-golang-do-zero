package main

import "fmt"

func main() {
	soma := somar(20, 6)

	fmt.Println(soma)

	var f = func(text string) {
		fmt.Println(text)
	}

	f("Função f")

	/* resultadoSoma, resultadoSubs := calculosMatematicos(10, 15)*/

	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println("soma:", resultadoSoma)
	/* fmt.Println("soma:", resultadoSubs) */

}

func somar(num1, num2 int8) int8 {
	return num1 + num2
}

/*Para declarar */

func calculosMatematicos(n1, n2 int16) (int16, int16) {

	soma := n1 + n2

	subs := n1 - n2

	return soma, subs
}
