package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type metaids map[string][]string

func (m metaids) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *metaids) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}

// Eroge represents a row in the `eroge` table and a singular file.
type Eroge struct {
	ID       int     `db:"id"`
	Filename string  `db:"fname"`
	Meta     metaids `db:"metaids"`
}

// User represents a row in the `users` table and a singular user.
type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (u User) String() string {
	return fmt.Sprintf("user %s", u.Username)
}
