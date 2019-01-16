package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SqlDb struct {
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

func Init(config string) (*SqlDb, error) {
	conn, err := sqlx.Open("postgres", config)
	if err != nil {
		return nil, err
	}
	s := &SqlDb{Conn: conn}
	if err := s.Conn.Ping(); err != nil {
		return nil, err
	}
	if err := s.createTables(); err != nil {
		return nil, err
	}
	if err := s.prepare(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SqlDb) createTables() error {
	createEro := `
		CREATE TABLE IF NOT EXISTS eroge (
		id SERIAL NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		title_jp TEXT NOT NULL,
		date TEXT NOT NULL,
		genres INTEGER[] NOT NULL,
		series INTEGER,
		ids TEXT[][],
		image TEXT NOT NULL,
		circle_id INTEGER NOT NULL
		);
	`
	rows, err := s.Conn.Query(createEro)
	if err != nil {
		return err
	}
	rows.Close()
	createCircles := `
		CREATE TABLE IF NOT EXISTS circles (
		id SERIAL NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		website TEXT NOT NULL
		);
	`
	rows, err = s.Conn.Query(createCircles)
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
	rows, err = s.Conn.Query(createSeries)
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
	rows, err = s.Conn.Query(createTags)
	if err != nil {
		return err
	}
	rows.Close()
	createEroTags := `
		CREATE TABLE IF NOT EXISTS ero_tags (
		ero_id INT,
		tag_name TEXT
		);
	`
	rows, err = s.Conn.Query(createEroTags)
	if err != nil {
		return err
	}
	rows.Close()
	rows, err = s.Conn.Query(createTags)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (s *SqlDb) prepare() error {
	var err error
	s.GetSomeEro, err = s.Conn.Preparex(
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
	s.GetAnEro, err = s.Conn.Preparex("SELECT * FROM eroge WHERE id=$1;")
	if err != nil {
		return err
	}
	s.GetEroTags, err = s.Conn.Preparex(
		`SELECT tag_name FROM eroge 
		INNER JOIN ero_tags ON 
		ero_tags.ero_id = eroge.id 
		WHERE eroge.id=$1`)
	if err != nil {
		return err
	}
	s.GetACircle, err = s.Conn.Preparex("SELECT * FROM circles WHERE id=$1;")
	if err != nil {
		return err
	}
	s.GetCircleEro, err = s.Conn.Preparex(
		`SELECT eroge.* FROM eroge 
		INNER JOIN circles 
		ON eroge.circle_name=circles.name 
		WHERE circles.id=$1;`)
	if err != nil {
		return err
	}
	s.IngestEro, err = s.Conn.PrepareNamed(
		`INSERT INTO eroge 
		(title, circle_name, dlsite_id, on_xdcc, on_hdd, in_torrent)
		VALUES 
		(:name, :circle, :id, :xdcc, :hdd, :torrent);`)
	if err != nil {
		return err
	}
	s.GetUnscrapedEro, err = s.Conn.Preparex(
		`SELECT dlsite_id 
		FROM eroge 
		WHERE scraped=False;`)
	if err != nil {
		return err
	}
	s.UpdateScrapedEro, err = s.Conn.Preparex(
		`UPDATE eroge
		SET scraped=True
		WHERE dlsite_id=$1;`)
	if err != nil {
		return err
	}

	return nil
}
