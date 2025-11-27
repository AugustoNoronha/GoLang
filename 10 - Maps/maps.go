package main

import "fmt"

func main() {
	fmt.Println("MAPS")
	//oque esta dentro dos [] são os tipos da chaves e o que esta fora o tipo do valor
	usuario := map[string]string{
		"nome":      "Augusto",
		"sobrenome": "Leite",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) //valores acessiveis desta forma

}
