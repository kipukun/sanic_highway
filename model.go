package main

import "github.com/lib/pq"

type Eroge struct {
	ID         int            `db:"id"`
	Title      string         `db:"title"`
	Date       string         `db:"date"`
	Images     pq.StringArray `db:"images"`
	CircleName string         `db:"circle_name"`
	DLsiteID   string         `db:"dlsite_id"`
	VNDBID     string         `db:"vndb_id"`
	MiscID     string         `db:"misc_id"`
	OnXDCC     bool           `db:"on_xdcc"`
	OnHDD      bool           `db:"on_hdd"`
	InTorrent  bool           `db:"in_torrent"`
	Circle     `db:"circle"`
}

type Circle struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Website string `db:"website"`
}
