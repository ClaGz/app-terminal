package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidores"
	flags := buildFlags("host", "devbook.com.br")

	app.Commands = []cli.Command{
		ipCommand(flags),
		servidoresCommand(flags),
	}
	return app
}

func buildFlags(name, value string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
}

func ipCommand(flags []cli.Flag) cli.Command {
	return cli.Command{
		Name:   "ip",
		Usage:  "Busca ips na net",
		Flags:  flags,
		Action: buscarIps,
	}
}
func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func servidoresCommand(flags []cli.Flag) cli.Command {
	return cli.Command{
		Name:   "servidores",
		Usage:  "Busca nomes dos servidores na net",
		Flags:  flags,
		Action: buscarServidores,
	}
}
func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
