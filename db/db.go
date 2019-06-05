package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Database holds the connection to the DB
// as well as prepared statements on that connection.
type Database struct {
	Conn                             *sqlx.DB
	Ero, Eros, IngestEro, InsertUser *sqlx.Stmt
}

type errExecer struct {
	c   *sqlx.DB
	err error
}

func (ee *errExecer) exec(stmt string) {
	if ee.err != nil {
		return
	}
	_, ee.err = ee.c.Exec(stmt)
}

func (ee *errExecer) prepare(prep **sqlx.Stmt, query string) {
	if ee.err != nil {
		return
	}
	*prep, ee.err = ee.c.Preparex(query)
}

// Init takes in a DB configuration string and returns a Database connection
// and nil, or nil and the error reported by prep functions otherwise.
func Init(config string) (*Database, error) {
	conn, err := sqlx.Open("postgres", config)
	if err != nil {
		return nil, err
	}
	d := &Database{Conn: conn}
	if err := d.Conn.Ping(); err != nil {
		return nil, err
	}
	if err := d.create(); err != nil {
		return nil, err
	}
	if err := d.prepare(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Database) create() error {
	ee := &errExecer{db: d.Conn}
	ee.exec(`CREATE TABLE IF NOT EXISTS eroge (
		    id SERIAL PRIMARY KEY,
		    fname text NOT NULL UNIQUE,
		    vndb_ids text[] NOT NULL,
		    dlsite_ids text[] NOT NULL );`)

	ee.exec("ALTER SEQUENCE eroge_id_seq RESTART WITH 1000;")

	ee.exec(`CREATE TABLE IF NOT EXISTS users (
		id uuid NOT NULL PRIMARY KEY,
		username text NOT NULL,
		password text NOT NULL );`)

	ee.exec(`CREATE TABLE IF NOT EXISTS sessions (
		id uuid NOT NULL,
		user_id uuid NOT NULL REFERENCES users(id),
		key text NOT NULL );`)

	if ee.err != nil {
		return ee.err
	}
	return nil
}

func (d *Database) prepare() error {
	ee := &errExecer{db: d.Conn}
	ee.prepare(&d.Eros, "SELECT * FROM eroge OFFSET $1 LIMIT 50;")
	ee.prepare(&d.Ero, "SELECT * FROM eroge WHERE id=$1;")
	ee.prepare(&d.IngestEro, `INSERT INTO eroge (fname) VALUES ($1)
		ON CONFLICT ON CONSTRAINT eroge_fname_key DO NOTHING;`)
	ee.prepare(&d.InsertUser, `INSERT INTO users (id, username, password) VALUES
				($1, $2, $3);`)
	if ee.err != nil {
		return ee.err
	}
	return nil
}
