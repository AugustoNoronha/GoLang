package main

import "fmt"

func f1() {
	fmt.Println("Formlar1")
}

func f2() {
	fmt.Println("Formlar2")
}

func main() {
	fmt.Println("defer")
	defer f1() // o defer da uma "segurada" no codigo
	// ele adia a execução da função
	f2()
}
