package db

import (
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func (d *Database) NewUser(uuid, user, hash string) error {
	_, err := d.InsertUser.Exec(uuid, user, hash)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code.Name() {
			case "unique_violation":
				return errors.New("username already exists")
			}
		}
		return err
	}
	return nil
}
