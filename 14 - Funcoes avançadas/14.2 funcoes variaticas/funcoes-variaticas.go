package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, item := range numeros {
		total += item
	}
	return total
}

func main() {
	soma := soma(10, 11, 12, 13, 14, 15)
	fmt.Println(soma)
}
