package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int16 = 100
	fmt.Println(num)

	var numGenerico int = 8888
	fmt.Println(numGenerico)

	var numero2 uint16 = 1000
	fmt.Println(numero2)

	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 33.00
	fmt.Println(numeroReal1)

	var str = "sajdasdas"
	fmt.Println(str)

	char := '%'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var boleano1 bool = false
	fmt.Println(boleano1)

	var err1 error = errors.New("erro interno")
	fmt.Println(err1)
}
