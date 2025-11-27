package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("USANDO STRUCT")
	var u usuario

	//imaginco que por caracteristas do GO não exita get e set
	u.nome = "Augusto Noronha"
	u.idade = 21
	fmt.Println(u)

	enderecoExempl := endereco{"Rua dos bobos", 0}

	u2 := usuario{"Nicolas Noronha", 24, enderecoExempl} //pode ser usado como construtor
	fmt.Println(u2)

	u3 := usuario{nome: "Jose Geraldo Leite"}
	fmt.Println(u3)

}
