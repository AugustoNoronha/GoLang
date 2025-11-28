package main

import "fmt"

func calculoMatematicos(n1 int, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calculoMatematicos(10, 20)
	fmt.Println(soma)
	fmt.Println(sub)
}
