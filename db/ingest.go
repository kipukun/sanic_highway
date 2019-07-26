package db

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/pkg/errors"
)

// Ingest takes in a Buffer and inserts each line into the DB.
// Returns nil on success, the error otherwise.
func (d *Database) Ingest(buf *bytes.Buffer) error {
	scanner := bufio.NewScanner(buf)
	tx, err := d.Conn.Beginx()
	if err != nil {
		return err
	}
	for scanner.Scan() {
		fname := strings.TrimSuffix(scanner.Text(), ".rar")
		_, err = tx.Stmtx(d.IngestEro).Exec(fname)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				err = errors.Wrap(err, "ingest: could not roll back")
				return err
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		err = errors.Wrap(err, "ingest: could not commit")
		return err
	}
	return nil
}
