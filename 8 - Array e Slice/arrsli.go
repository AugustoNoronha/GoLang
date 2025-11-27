package main

import "fmt"

func main() {
	fmt.Println("ARRAY E SLICE")

	var arr [5]int //é necessario especificar o tamanho do array (curioso para o equivalente a List)
	fmt.Println(arr)

	for i := 0; i < 5; i++ {
		arr[i] = i + 10
		fmt.Println(arr[i])
	}

	fmt.Println("=================================")

	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		slice[i] = i + 30
		fmt.Println(slice[i])
	}

	//para adcionar novos itens no slice devemos usar o append
	slice = append(slice, 19)
	fmt.Println("======")
	fmt.Println(slice)

	//podemos usar a função make() para criar slices
	slice2 := make([]int, 10, 15)
	fmt.Println("======")
	fmt.Println(slice2)
	fmt.Println(len(slice2))   //retorna tamanho
	fmt.Println(cap((slice2))) //retnar capacidade maxima
	//caso ultrapace a capciade de um slice ele dobra a capcidade, desta forma nunca ficando sem espaço
	//tal qua um array
	//o slice pode ser utilizado com uma Lista na teroia

}
