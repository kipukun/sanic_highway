package db

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/lib/pq"
)

// root is just for holding the parsed XML.
type root struct {
	XMLName     xml.Name `xml:"root"`
	Text        string   `xml:",chardata"`
	Erofilelist struct {
		Text    string `xml:",chardata"`
		Erofile []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name"`
			Tags string `xml:"tags"`
			Ids  struct {
				Text string   `xml:",chardata"`
				ID   []string `xml:"id"`
			} `xml:"ids"`
			Onxdcc    int `xml:"onxdcc"`
			Onhdd     int `xml:"onhdd"`
			Intorrent int `xml:"intorrent"`
		} `xml:"erofile"`
	} `xml:"erofilelist"`
}

// ero contains Ero ready for insertion into the DB.
type ero struct {
	Name      string
	DLsiteIDs pq.StringArray
	VNDBIDs   pq.StringArray
	Circle    string
	XDCC      bool
	HDD       bool
	Torrent   bool
}

// parses XML and returns an ErofileList and nil, or an empty object and the error.
func parse(file string) (root, error) {
	var rt root
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ingest(): error reading file:", err.Error())
		return rt, err
	}

	fmt.Println("opened", file)

	err = xml.Unmarshal(f, &rt)
	if err != nil {
		fmt.Println("ingest(): error unmarshalling file:", err.Error())
		return rt, err
	}

	return rt, nil
}

// processes an ErofileList for insertion into the database.
func process(rt root) []ero {
	var pel []ero
	for _, entry := range rt.Erofilelist.Erofile {
		var e ero
		var dlid []string
		var vnid []string
		f := strings.TrimSuffix(entry.Name, ".rar")
		split := strings.Split(f, "]")
		c := strings.Trim(split[0], "[] ")
		n := strings.Trim(split[1], "[] ")
		for _, id := range entry.Ids.ID {
			if id != "" {
				switch id[:2] {
				case "RE", "RJ", "RG", "VJ", "VG":
					dlid = append(dlid, id)
				}
				switch id[:1] {
				case "v", "r", "p":
					vnid = append(vnid, id)
				}
			}

		}
		if entry.Onhdd == 1 {
			e.HDD = true
		}
		if entry.Onxdcc == 1 {
			e.XDCC = true
		}
		if entry.Intorrent == 1 {
			e.Torrent = true
		}

		e.Name = n
		e.DLsiteIDs = pq.StringArray(dlid)
		e.VNDBIDs = pq.StringArray(vnid)
		e.Circle = c
		pel = append(pel, e)
	}
	return pel
}

// puts the parsed XML inside the database.
func (d *Database) IngestXML(file string) error {
	var err error
	rt, err := parse(file)
	if err != nil {
		fmt.Println("IngestXML: error parsing file", err.Error())
		return err
	}
	pel := process(rt)
	for _, e := range pel {
		_, err = d.IngestEro.Exec(e)
		if err != nil {
			fmt.Println("IngestXML: error exec'ing", err.Error())
			return err
		}
		_, err = d.UpdateCircles.Exec(e.Circle)
		if err != nil {
			fmt.Println("UpdateCircles: error putting:", err.Error())
			return err
		}
	}
	return nil
}
