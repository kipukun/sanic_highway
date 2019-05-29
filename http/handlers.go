package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kipukun/sanic_highway/model"
	"github.com/kipukun/sanic_highway/templates"
)

func (s *Server) indexHandler(index bool) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var erogeList []model.Eroge
		var pg int
		var err error
		var pn [5]int

		if index {
			pg = 0
		} else {
			vars := mux.Vars(r)
			pg, err = strconv.Atoi(vars["page"])
			if err != nil {
				return
			}
		}

		err = s.db.GetSomeEro.Select(&erogeList, pg*50)
		if err != nil {
			log.Println("indexHandler: GetSomeEro query failed", err.Error())
			return
		}

		// generate numbers for pagination
		switch pg {
		case 0, 1, 2, 3:
			for i, _ := range pn {
				pn[i] = i
			}
		default:
			for i, j := int(pg)-2, 0; j < 5; i, j = i+1, j+1 {
				pn[j] = i
			}
		}
		p := &templates.IndexPage{
			erogeList,
			pn,
			int(pg),
			int(pg) - 1,
			int(pg) + 1,
		}

		templates.WritePageTemplate(w, p)

		return
	})
}

func (s *Server) aboutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := &templates.AboutPage{}
		templates.WritePageTemplate(w, p)
		return
	})
}

func (s *Server) stopHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := &templates.StopPage{}
		templates.WritePageTemplate(w, p)
		return
	})
}

func (s *Server) eroHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var eroge model.Eroge
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Println("strconv:", err.Error())
			return

		}
		err = s.db.GetAnEro.Get(&eroge, id)
		if err != nil {
			p := &templates.StopPage{}
			templates.WritePageTemplate(w, p)
			log.Println("eroHandler: GetAnEro query failed", err.Error())
			return
		}
		p := &templates.ErogePage{
			Ero: eroge,
		}

		templates.WritePageTemplate(w, p)

		return
	})
}

/*
func (s *Server) circleHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var circle model.Circle
		var ero []model.Eroge
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Println("strconv:", err.Error())
			return

		}

		err = s.db.GetACircle.Get(&circle, id)
		if err != nil {
			p := &templates.StopPage{}
			templates.WritePageTemplate(w, p)
			log.Println("circleHandler: GetACircle query failed", err.Error())
			return
		}
		err = s.db.GetCircleEro.Select(&ero, id)
		if err != nil {
			p := &templates.StopPage{}
			templates.WritePageTemplate(w, p)
			log.Println("circleHandler: GetCircleEro query failed", err.Error())
			return
		}

		p := &templates.CirclePage{
			Circle: circle,
			Ero:    ero,
		}

		templates.WritePageTemplate(w, p)

		return
	})
}
*/
