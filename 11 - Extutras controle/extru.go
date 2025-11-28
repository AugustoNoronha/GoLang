package main

import "fmt"

func main() {
	fmt.Println(("Extrturas de controle"))

	numero := 10
	if numero > 15 {
		fmt.Println("Sim")
	} else {
		fmt.Println("Não")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("é maior que 0")
	}

	for indice, item := range array { //isso é basicamnete o forEacth
		println("indeice", indice)
		println("item", item)

	}

}
