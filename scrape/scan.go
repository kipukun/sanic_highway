package scrape

import (
	"errors"
	"fmt"
	"time"

	"github.com/kipukun/sanic_highway/db"
)

// scan takes in a SqlDb, and channels for both IDs to process and error to be sent up the chain.
func scan(d *db.SqlDb, ids chan string, errCh chan error) {
	rows := make([]string, 1)
	err := d.GetUnscrapedEro.Select(&rows)
	if err != nil {
		errCh <- err
	}
	for _, id := range rows {
		ids <- id
		time.Sleep(10 * time.Second)
	}
	close(ids)
}

func Start(d *db.SqlDb, errCh chan error) {
	var err error
	ids := make(chan string)
	go scan(d, ids, errCh)
	for {
		id, ok := <-ids
		if ok == false {
			err = errors.New("Start(): Error getting ID on channel.")
			errCh <- err
		}
		if id != "" {
			fmt.Println("[scrape] Scraping: ", id)
			_, err := d.UpdateScrapedEro.Exec(id)
			if err != nil {
				errCh <- err
			}
		}
	}
}
