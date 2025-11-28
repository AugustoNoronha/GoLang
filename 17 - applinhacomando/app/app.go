package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// gerar vai restoarnar a aplicação de linha de comando pronta rpa ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "APP de linha de comnado"
	app.Usage = "Encontar ip e servidor de sites"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de ips na internte",
			Flags:  flags,
			Action: buscarIpds,
		}, {
			Name:   "servidores",
			Usage:  "Buscar servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIpds(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}
