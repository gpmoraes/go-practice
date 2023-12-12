package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	// Creates a new instance of the CLI application
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Get IPs and Server's names in the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}
	// Defines two commands: "ip" and "servers"
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Get IPs from internet URLS",
			Flags:  flags,
			Action: getIps,
		},
		{
			Name:   "servers",
			Usage:  "Get the server's name",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

// Function to fetch and print IPs for a given host
func getIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Function to fetch and print server names for a given host
func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
