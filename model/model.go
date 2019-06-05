package model

// Eroge represents a row in the `eroge` table and a singular file.
type Eroge struct {
	ID       int      `db:"id"`
	Filename string   `db:"fname"`
	DLsite   []string `db:"dlsite_ids"`
	VNDB     []string `db:"vndb_ids"`
}
