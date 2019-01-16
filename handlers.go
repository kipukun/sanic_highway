package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/kipukun/sanic_highway/db"
)

var templates = template.Must(template.ParseGlob("tmpl/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, d interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func parseParam(w http.ResponseWriter, r *http.Request, param string) (int64, error) {
	q, ok := r.URL.Query()[param]
	if !ok || len(q) < 1 {
		http.Error(w, "Invalid query.", http.StatusBadRequest)
		return 0, errors.New("[!] parseID: error parsing param from query")
	}
	id, err := strconv.ParseInt(q[0], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0, err
	}

	return id, nil

}

func indexHandler(db *db.SqlDb) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var erogeList []Eroge
		var id int64
		var err error
		var pn [5]int
		// check if the param is there
		// if not, assume we're on the index page and return first 10
		_, ok := r.URL.Query()["p"]
		if !ok {
			id = 0
		} else {
			id, err = parseParam(w, r, "p")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		err = db.GetSomeEro.Select(&erogeList, id*50)
		if err != nil {
			fmt.Printf("[!] GetAllEro query failed!\n %s", err)
			return
		}

		// generate numbers for pagination
		switch id {
		case 0, 1, 2, 3:
			for i, _ := range pn {
				pn[i] = i
			}
		default:
			for i, j := int(id)-2, 0; j < 5; i, j = i+1, j+1 {
				pn[j] = i
			}
		}

		data := struct {
			ErogeList  []Eroge
			Pagination [5]int
			Current    int
			Prev       int
			Next       int
		}{
			erogeList,
			pn,
			int(id),
			int(id) - 1,
			int(id) + 1,
		}

		renderTemplate(w, "index", data)
		return
	})
}

func eroHandler(db *db.SqlDb) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var eroge []Eroge
		var tags []string
		id, err := parseParam(w, r, "id")
		if err != nil {
			return
		}
		err = db.GetAnEro.Select(&eroge, id)
		if err != nil {
			fmt.Printf("[!] GetAnEro query failed!\n %s", err)
		}
		err = db.GetEroTags.Select(&tags, id)
		if err != nil {
			fmt.Printf("[!] GetEroTags query failed!\n %s", err)
		}

		data := struct {
			Ero  []Eroge
			Tags []string
		}{
			eroge,
			tags,
		}

		renderTemplate(w, "ero", data)
		return
	})
}

func circleHandler(db *db.SqlDb) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var circle []Circle
		var ero []Eroge
		id, err := parseParam(w, r, "id")
		if err != nil {
			return
		}
		err = db.GetACircle.Select(&circle, id)
		if err != nil {
			fmt.Printf("[!] GetAnEro query failed!\n %s", err)
		}
		err = db.GetCircleEro.Select(&ero, id)
		if err != nil {
			fmt.Printf("[!] GetCircleEro query failed!\n %s", err)
		}

		data := struct {
			Circle []Circle
			Ero    []Eroge
		}{
			circle,
			ero,
		}

		renderTemplate(w, "circle", data)

	})
}
