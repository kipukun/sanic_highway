package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func main() {
	/* db
	fmt.Println("[*] initializing db...")
	d, err := db.Init("postgres://postgres:cock@localhost/pqgodb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()
	*/

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&scraperCmd{}, "")
	subcommands.Register(&ingestCmd{}, "")
	flag.Parse()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))

	/*
		fmt.Println("[*] starting up http...")

		srv := http.Init(d)
		srv.Start() */

}
