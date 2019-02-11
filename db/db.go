package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Database holds the connection to the DB
// as well as prepared statements on that connection.
type Database struct {
	Conn *sqlx.DB

	GetAnEro         *sqlx.Stmt
	GetSomeEro       *sqlx.Stmt
	GetEroTags       *sqlx.Stmt
	GetACircle       *sqlx.Stmt
	GetCircleEro     *sqlx.Stmt
	IngestEro        *sqlx.NamedStmt
	GetUnscrapedEro  *sqlx.Stmt
	UpdateScrapedEro *sqlx.Stmt
}

func Init(config string) (*Database, error) {
	conn, err := sqlx.Open("postgres", config)
	if err != nil {
		return nil, err
	}
	d := &Database{Conn: conn}
	if err := d.Conn.Ping(); err != nil {
		return nil, err
	}
	if err := d.createTables(); err != nil {
		return nil, err
	}
	if err := d.prepare(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Database) createTables() error {
	createEro := `
		CREATE TABLE IF NOT EXISTS eroge (
		    id integer NOT NULL,
		    title text NOT NULL,
		    circle_name text,
		    dlsite_ids text[],
		    vndb_ids text[],
		    misc_ids text[],
		    date text DEFAULT '11/2/1971'::text NOT NULL,
		    on_xdcc boolean DEFAULT false NOT NULL,
		    on_hdd boolean DEFAULT false NOT NULL,
		    in_torrent boolean DEFAULT false NOT NULL,
		    images text[] DEFAULT '{https://via.placeholder.com/400x400.png?text=cum,https://via.placeholder.com/400x400.png?text=cum,https://via.placeholder.com/400x400.png?text=cum}'::text[] NOT NULL,
		    scraped boolean DEFAULT false NOT NULL
		);
	`
	rows, err := d.Conn.Query(createEro)
	if err != nil {
		return err
	}
	rows.Close()
	createCircles := `
		CREATE TABLE IF NOT EXISTS circles (
		id SERIAL NOT NULL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		website TEXT NOT NULL
		);
	`
	rows, err = d.Conn.Query(createCircles)
	if err != nil {
		return err
	}
	rows.Close()
	createSeries := `
		CREATE TABLE IF NOT EXISTS series (
		id SERIAL NOT NULL PRIMARY KEY,
		name TEXT NOT NULL
		);
	`
	rows, err = d.Conn.Query(createSeries)
	if err != nil {
		return err
	}
	rows.Close()
	createTags := `
		CREATE TABLE IF NOT EXISTS tags (
		id SERIAL NOT NULL PRIMARY KEY,
		name TEXT NOT NULL
		);
	`
	rows, err = d.Conn.Query(createTags)
	if err != nil {
		return err
	}
	rows.Close()
	createEroTags := `
		CREATE TABLE IF NOT EXISTS ero_tags (
		ero_id INT NOT NULL,
		tag_name TEXT NOT NULL 
		);
	`
	rows, err = d.Conn.Query(createEroTags)
	if err != nil {
		return err
	}
	rows.Close()
	rows, err = d.Conn.Query(createTags)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (d *Database) prepare() error {
	var err error
	d.GetSomeEro, err = d.Conn.Preparex(
		`SELECT eroge.*,
			circles.id AS "circle.id"
		FROM eroge 
		INNER JOIN circles 
		ON eroge.circle_name=circles.name
		OFFSET $1
		LIMIT 50;`)
	if err != nil {
		return err
	}
	d.GetAnEro, err = d.Conn.Preparex("SELECT * FROM eroge WHERE id=$1;")
	if err != nil {
		return err
	}
	d.GetEroTags, err = d.Conn.Preparex(
		`SELECT tag_name FROM eroge 
		INNER JOIN ero_tags ON 
		ero_tags.ero_id = eroge.id 
		WHERE eroge.id=$1`)
	if err != nil {
		return err
	}
	d.GetACircle, err = d.Conn.Preparex("SELECT * FROM circles WHERE id=$1;")
	if err != nil {
		return err
	}
	d.GetCircleEro, err = d.Conn.Preparex(
		`SELECT eroge.* FROM eroge 
		INNER JOIN circles 
		ON eroge.circle_name=circles.name 
		WHERE circles.id=$1;`)
	if err != nil {
		return err
	}
	d.IngestEro, err = d.Conn.PrepareNamed(
		`INSERT INTO eroge 
		(title, circle_name, dlsite_ids, vndb_ids, on_xdcc, on_hdd, in_torrent)
		VALUES 
		(:name, :circle, :dlsiteids, :vndbids, :xdcc, :hdd, :torrent);`)
	if err != nil {
		return err
	}
	d.GetUnscrapedEro, err = d.Conn.Preparex(
		`SELECT dlsite_ids 
		FROM eroge 
		WHERE scraped=False;`)
	if err != nil {
		return err
	}
	d.UpdateScrapedEro, err = d.Conn.Preparex(
		`UPDATE eroge
		SET scraped=True
		WHERE dlsite_ids=$1;`)
	if err != nil {
		return err
	}

	return nil
}
