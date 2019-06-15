package model

import (
	"fmt"

	"github.com/lib/pq"
)

// Eroge represents a row in the `eroge` table and a singular file.
type Eroge struct {
	ID       int            `db:"id"`
	Filename string         `db:"fname"`
	DLsite   pq.StringArray `db:"dlsite_ids"`
	VNDB     pq.StringArray `db:"vndb_ids"`
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
