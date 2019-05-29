package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kipukun/sanic_highway/db"
	"github.com/kipukun/sanic_highway/http"
)

func main() {
	fmt.Println("[*] initializing db...")
	d, err := db.Init("postgres://postgres:cock@localhost/pqgodb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	var i = flag.String("ingest", "NULL", "file to ingest")
	flag.Parse()

	fmt.Println(*i)
	if *i != "NULL" {
		f, err := os.Open(*i)
		if err != nil {
			panic(err)
		}
		db.Parse(f, d)
	}

	fmt.Println("[*] starting up http...")

	srv := http.Init(d)
	srv.Start()

}
