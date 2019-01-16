package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kipukun/sanic_highway/db"
	"github.com/kipukun/sanic_highway/scrape"
)

func load(name string) string {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	return string(f)
}

func static(page string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, page)
		return
	})
}

func main() {
	ingestPtr := flag.String("ingest", "NULL", "file to ingest")
	flag.Parse()

	// db
	fmt.Println("[*] initializing db...")
	db, err := db.Init("postgres://pqgo:cock@localhost/pqgodb?sslmode=require")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Conn.Close()

	// scraper
	errCh := make(chan error)

	go func() {
		scrape.Start(db, errCh)
		for {
			select {
			case <-errCh:
				log.Fatalln(<-errCh)
			}
		}
	}()

	if *ingestPtr != "NULL" {
		fmt.Printf("[*] ingest mode detected\n[*] opening file " + *ingestPtr + "\n")
		el, err := ingest(*ingestPtr)
		if err != nil {
			return
		}
		pel := process(el)
		fmt.Println("[*] inserting into db...")
		err = putXML(pel, db)
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	}

	fmt.Println("[*] loading static...")
	about := load("tmpl/about.html")

	fmt.Println("[*] setting up http...")
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/", indexHandler(db))
	http.Handle("/ero", eroHandler(db))
	http.Handle("/circle", circleHandler(db))
	http.Handle("/about", static(about))
	log.Fatal(http.ListenAndServe(":1337", nil))
}
