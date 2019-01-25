package model

import "github.com/lib/pq"

type Eroge struct {
	ID         int            `db:"id"`
	Title      string         `db:"title"`
	Date       string         `db:"date"`
	Images     pq.StringArray `db:"images"`
	CircleName string         `db:"circle_name"`
	DLsiteIDs  pq.StringArray `db:"dlsite_ids"`
	VNDBIDs    pq.StringArray `db:"vndb_ids"`
	MiscIDs    pq.StringArray `db:"misc_ids"`
	OnXDCC     bool           `db:"on_xdcc"`
	OnHDD      bool           `db:"on_hdd"`
	InTorrent  bool           `db:"in_torrent"`
	Scraped    bool           `db:"scraped"`
	Circle     `db:"circle"`
}

type Circle struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Website string `db:"website"`
}
