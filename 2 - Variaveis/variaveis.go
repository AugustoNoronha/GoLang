package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Valor variavel 2 "
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "vareavel3"
		variavel4 string = "vareavel3"
	)
	fmt.Println(variavel3, variavel4)
	variavel5, variavel6 := "varavel5", "variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 int = 80
	fmt.Println(constante1)
}
