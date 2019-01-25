package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kipukun/sanic_highway/db"
	"github.com/kipukun/sanic_highway/model"
	"github.com/kipukun/sanic_highway/templates"
)

// Server holds the DB connection, configured routes for the http package
// as well as the parsed templates.
type Server struct {
	db     *db.Database
	routes *mux.Router
}

func Init(path string, d *db.Database) (*Server, error) {
	r := mux.NewRouter()

	srv := &Server{
		db:     d,
		routes: r,
	}

	return srv, nil
}

func (s *Server) Start() {
	s.routes.Handle("/", s.indexHandler(true))
	s.routes.Handle("/page/{page}", s.indexHandler(false))
	// heh
	ass := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	s.routes.PathPrefix("/assets/").Handler(ass)
	s.routes.Handle("/ero/{id}", s.eroHandler())
	s.routes.Handle("/circle/{id}", s.circleHandler())

	log.Fatal(http.ListenAndServe(":1337", s.routes))
}

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
			fmt.Printf("[!] GetAllEro query failed!\n %s", err)
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

func (s *Server) eroHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var eroge []model.Eroge
		var tags []string
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		err = s.db.GetAnEro.Select(&eroge, id)
		if err != nil {
			fmt.Printf("[!] GetAnEro query failed!\n %s", err)
		}
		err = s.db.GetEroTags.Select(&tags, id)
		if err != nil {
			fmt.Printf("[!] GetEroTags query failed!\n %s", err)
		}

		p := &templates.ErogePage{
			Ero:  eroge,
			Tags: tags,
		}

		templates.WritePageTemplate(w, p)

		return
	})
}

func (s *Server) circleHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var circle []model.Circle
		var ero []model.Eroge
		id := 0
		err := s.db.GetACircle.Select(&circle, id)
		if err != nil {
			fmt.Printf("[!] GetAnEro query failed!\n %s", err)
		}
		err = s.db.GetCircleEro.Select(&ero, id)
		if err != nil {
			fmt.Printf("[!] GetCircleEro query failed!\n %s", err)
		}

		p := &templates.CirclePage{
			Circle: circle,
			Ero:    ero,
		}

		templates.WritePageTemplate(w, p)

		return
	})
}
