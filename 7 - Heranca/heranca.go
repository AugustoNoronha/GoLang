package main

import "fmt"

type pessoa struct {
	nome   string
	idiade uint8
}

type estudate struct {
	pessoa
	curso     string
	matricula uint32
}

func main() {
	fmt.Println("HERANÇA SO QUE NAO ") //so que nao mesmo
	p := pessoa{"tulho", 18}
	e := estudate{p, "Ciencias da Comuputação", 783624}

	fmt.Println(e)
	fmt.Println(e.nome)
	fmt.Println(e.idiade)
	fmt.Println(e.curso)
	fmt.Println(e.matricula)

}
