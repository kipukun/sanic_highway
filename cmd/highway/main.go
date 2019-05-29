package main

import (
	"fmt"
	"log"

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

	fmt.Println("[*] starting up http...")

	srv := http.Init(d)
	srv.Start()

}
