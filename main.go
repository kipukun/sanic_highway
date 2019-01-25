//go:generate qtc -dir=templates
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kipukun/sanic_highway/db"
	"github.com/kipukun/sanic_highway/http"
	"github.com/kipukun/sanic_highway/scrape"
)

func main() {
	ingestPtr := flag.String("ingest", "NULL", "file to ingest")
	scrapePtr := flag.Bool("scrape", false, "enable scraping mode")
	flag.Parse()

	// db
	fmt.Println("[*] initializing db...")
	d, err := db.Init("postgres://pqgo:cock@localhost/pqgodb?sslmode=require")
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	if *scrapePtr {
		// scraper
		errCh := make(chan error)
		go func() {
			scrape.Start(d, errCh)
			for {
				select {
				case <-errCh:
					log.Fatalln(<-errCh)
				}
			}
		}()
	}

	if *ingestPtr != "NULL" {
		fmt.Println("[*] starting ingest...")
		err := d.IngestXML(*ingestPtr)
		if err != nil {
			fmt.Println("[!] IngestXML failed:", err.Error())
			return
		}
		return
	}

	fmt.Println("[*] setting up http...")

	srv, err := http.Init("tmpl/", d)
	if err != nil {
		panic(err)
	}

	srv.Start()

}
