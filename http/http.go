package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kipukun/sanic_highway/db"
)

// Server holds the DB connection, configured routes for the http package
// as well as the parsed templates.
type Server struct {
	db     *db.Database
	routes *mux.Router
}

func Init(path string, d *db.Database) *Server {
	r := mux.NewRouter()

	srv := &Server{
		db:     d,
		routes: r,
	}

	return srv
}

// Start sets up routes on the mux and starts the HTTP server.
func (s *Server) Start() {
	s.routes.Handle("/", s.indexHandler(true))
	s.routes.Handle("/about", s.aboutHandler())
	s.routes.Handle("/page/{page}", s.indexHandler(false))
	// heh
	ass := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	s.routes.PathPrefix("/assets/").Handler(ass)
	s.routes.Handle("/ero/{id}", s.eroHandler())
	s.routes.Handle("/circle/{id}", s.circleHandler())

	s.routes.NotFoundHandler = s.stopHandler()

	log.Fatal(http.ListenAndServe(":1337", s.routes))
}
