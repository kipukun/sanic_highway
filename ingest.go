package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/kipukun/sanic_highway/db"
)

// ErofileList is just for holding the XML.
type erofileList struct {
	Erofile []struct {
		Name    string   `xml:"name"`
		IDs     []string `xml:"ids>id"`
		XDCC    int      `xml:"onxdcc"`
		HDD     int      `xml:"onhdd"`
		Torrent int      `xml:"intorrent"`
	} `xml:"erofilelist>erofile"`
}

// ProccesedEro contains Ero ready for insertion into the DB.
type ero struct {
	Name    string
	ID      string
	Circle  string
	XDCC    bool
	HDD     bool
	Torrent bool
}

// parses XML and returns an ErofileList and nil, or an empty object and the error.
func ingest(file string) (erofileList, error) {
	var el erofileList
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ingest(): error reading file:" + err.Error())
		return el, err
	}

	fmt.Println("opened " + file)

	err = xml.Unmarshal(f, &el)
	if err != nil {
		fmt.Println("ingest(): error unmarshalling file:" + err.Error())
		return el, err
	}

	return el, nil
}

// processes an ErofileList for insertion into the database.
func process(el erofileList) []ero {
	var pel []ero
	for _, entry := range el.Erofile {
		var e ero
		var i string
		f := strings.TrimSuffix(entry.Name, ".rar")
		split := strings.Split(f, "]")
		c := strings.Trim(split[0], "[] ")
		n := strings.Trim(split[1], "[] ")
		for _, id := range entry.IDs {
			switch id[:2] {
			case "RE", "RJ":
				i = id
			}

		}
		if entry.HDD == 1 {
			e.HDD = true
		}
		if entry.XDCC == 1 {
			e.XDCC = true
		}
		if entry.Torrent == 1 {
			e.Torrent = true
		}

		e.Name = n
		e.ID = i
		e.Circle = c
		pel = append(pel, e)
	}
	return pel
}

// puts the parsed XML inside the database.
func putXML(pel []ero, s *db.SqlDb) error {
	var err error
	for _, e := range pel {
		_, err = s.IngestEro.Exec(e)
		if err != nil {
			fmt.Println("putXML(): error exec'ing " + err.Error())
		}
	}
	if err != nil {
		fmt.Println("putXML(): error commiting tx " + err.Error())
		return err
	}
	return nil
}
