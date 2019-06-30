package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Database holds the connection to the DB
// as well as prepared statements on that connection.
type Database struct {
	Conn                                            *sqlx.DB
	Ero, Eros, IngestEro, InsertUser, CreateSession *sqlx.Stmt
	CreateMeta, RemoveMeta, UpdateMeta, DeleteMeta  *sqlx.Stmt
	User, Lookup                                    *sqlx.Stmt
}

type errExecer struct {
	c   *sqlx.DB
	err error
}

func (ee *errExecer) exec(stmt string) {
	if ee.err != nil {
		return
	}
	_, err := ee.c.Exec(stmt)
	ee.err = errors.Wrap(err, stmt)
}

func (ee *errExecer) prepare(prep **sqlx.Stmt, query string) {
	var err error
	if ee.err != nil {
		return
	}
	*prep, err = ee.c.Preparex(query)
	ee.err = errors.Wrap(err, query)
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
	ee := &errExecer{c: d.Conn}
	ee.exec(`CREATE TABLE IF NOT EXISTS eroge (
			id SERIAL PRIMARY KEY,
			fname text NOT NULL UNIQUE, 
			metaids jsonb DEFAULT '{}');`)

	ee.exec("ALTER SEQUENCE eroge_id_seq RESTART WITH 1000;")

	ee.exec(`CREATE TABLE IF NOT EXISTS users (
			id uuid NOT NULL PRIMARY KEY,
			username text UNIQUE NOT NULL,
			password text NOT NULL );`)

	ee.exec(`CREATE TABLE IF NOT EXISTS sessions (
			id uuid NOT NULL,
			user_id uuid UNIQUE NOT NULL REFERENCES users(id) );`)

	if ee.err != nil {
		return ee.err
	}
	return nil
}

func (d *Database) prepare() error {
	ee := &errExecer{c: d.Conn}
	ee.prepare(&d.Eros, "SELECT * FROM eroge ORDER BY id ASC OFFSET $1 LIMIT $2;")
	ee.prepare(&d.Ero, `SELECT * FROM eroge WHERE id = $1`)
	ee.prepare(&d.IngestEro, `INSERT INTO eroge (fname) VALUES ($1)
		ON CONFLICT ON CONSTRAINT eroge_fname_key DO NOTHING;`)
	ee.prepare(&d.InsertUser, `INSERT INTO users (id, username, password) VALUES
				($1, $2, $3);`)
	ee.prepare(&d.CreateSession, `INSERT INTO sessions (id, user_id) 
				VALUES ($1, $2)
				ON CONFLICT ON CONSTRAINT sessions_user_id_key
				DO UPDATE SET id = $1;`)
	ee.prepare(&d.CreateMeta, `UPDATE eroge SET metaids =
	metaids || jsonb_build_object($1::text, $2::jsonb) WHERE id = $3;`)
	ee.prepare(&d.RemoveMeta, `UPDATE eroge SET metaids =
			jsonb_set(metaids, $1::text[], (metaids->$2::text) - -1)
			WHERE id = $3;`)
	ee.prepare(&d.UpdateMeta, `UPDATE eroge SET metaids =
	jsonb_insert(metaids, $1::text[], $2::jsonb, true) WHERE id = $3;`)
	ee.prepare(&d.DeleteMeta, `UPDATE eroge SET metaids = metaids - $1::text
			WHERE id = $2;`)
	ee.prepare(&d.User, `SELECT * FROM users WHERE username = $1 LIMIT 1;`)
	ee.prepare(&d.Lookup, `SELECT users.* FROM users
			INNER JOIN sessions 
			ON sessions.user_id = users.id
			AND sessions.id = $1;`)
	if ee.err != nil {
		return ee.err
	}
	return nil
}
