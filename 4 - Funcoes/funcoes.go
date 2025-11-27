package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosmatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(8, 9)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("TEXTO F1")
	fmt.Println(resultado)

	resultadosoma, resultadosub := calculosmatematicos(5, 6)
	fmt.Println(resultadosoma, resultadosub)
}
