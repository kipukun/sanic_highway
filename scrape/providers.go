package scrape

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gocolly/colly"
)

type Chobit struct {
	Works []struct {
		EmbedURL string `json:"embed_url"`
	} `json:"works"`
}

type VNDB struct {
}

func dlsite(id string) {
	c := colly.NewCollector()

	// finds the video tag when visiting the embed URL
	c.OnHTML("meta[itemprop=contentUrl]", func(e *colly.HTMLElement) {
		c.Visit(e.Attr("content"))
	})

	// finds images on the main page
	c.OnHTML("div[data-src]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("data-src"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		if _, err := os.Stat("./data/dlsite/" + id); os.IsNotExist(err) {
			os.Mkdir("./data/dlsite/"+id, 0750)
		}

		// download an image/video if we get it
		switch r.Headers.Get("Content-Type") {
		case "image/jpeg", "video/mp4":
			r.Save("./data/dlsite/" + id + "/" + r.FileName())
		case "text/javascript":
			container := new(Chobit)
			trim := bytes.TrimPrefix(r.Body, []byte("response("))
			data := bytes.TrimRight(trim, ")")
			json.Unmarshal(data, &container)
			c.Visit(container.Works[0].EmbedURL)
		case "text/html; charset=UTF-8":
			fmt.Println("Saving content of " + id)

			err := ioutil.WriteFile("./data/dlsite/"+id+"/body.html", r.Body, 0750)
			if err != nil {
				fmt.Println("[*] Error writing file for " + r.Request.URL.RequestURI())
				fmt.Println(err.Error())
			}
			return

		}

	})

	c.Visit("https://www.dlsite.com/maniax/work/=/product_id/" + id + ".html")
	c.Visit("https://chobit.cc/api/v1/dlsite/embed?workno=" + id)
}

func vndb(id string) {
	c := colly.NewCollector()

	c.OnHTML("a.scrlnk", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// save the HTML content of the page
	c.OnResponse(func(r *colly.Response) {

		if _, err := os.Stat("./data/vndb/v" + id); os.IsNotExist(err) {
			os.Mkdir("./data/vndb/v"+id, 0750)
		}

		switch r.Headers.Get("content-type") {
		case "text/html; charset=UTF-8":
			fmt.Println("Saving content of v" + id)

			err := ioutil.WriteFile("./data/vndb/v"+id+"/body.html", r.Body, 0750)
			if err != nil {
				fmt.Println("[*] Error writing file for " + r.Request.URL.RequestURI())
				fmt.Println(err.Error())
			}
			return
		case "image/jpeg":
			fmt.Println("Saving " + r.FileName())
			r.Save("./data/vndb/v" + id + "/" + r.FileName())
		}

	})

	c.Visit("https://vndb.org/v" + id)
}
