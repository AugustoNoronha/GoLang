package main

import (
	"applinhacomando/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("EXECUTANDO MAISN")
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
