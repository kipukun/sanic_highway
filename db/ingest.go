package db

import (
	"bufio"
	"os"
	"strings"
)

// Parse takes in pointers to a file and database, and inserts
// each line into the database. Returns an error if encountered.
func Parse(f *os.File, d *Database) error {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fname := strings.TrimSuffix(scanner.Text(), ".rar")
		_, err := d.IngestEro.Exec(fname)
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
