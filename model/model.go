package model

type Eroge struct {
	ID       int      `db:"id"`
	Filename string   `db:"fname"`
	DLsite   []string `db:"dlsite_ids"`
	VNDB     []string `db:"vndb_ids"`
}

type Circle struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
