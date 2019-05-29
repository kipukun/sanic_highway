package model

type Eroge struct {
	ID       int    `db:"id"`
	Filename string `db:"fname"`
}

type Circle struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
