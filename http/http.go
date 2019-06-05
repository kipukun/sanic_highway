package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/kipukun/sanic_highway/db"
)

// Server holds the DB connection, and configured routes for the http package.
type Server struct {
	DB     *db.Database
	routes *mux.Router
}

// Init takes in a Database and returns a new Server intialized with the
// default mux.Router and the database connection.
func Init(d *db.Database) *Server {
	r := mux.NewRouter()
	srv := &Server{
		DB:     d,
		routes: r,
	}
	return srv
}

// Start sets up routes on the mux and starts the HTTP server.
func (s *Server) Start() {
	s.routes.Handle("/", Handler{s, getIndex})
	s.routes.Handle("/about", Handler{s, getAbout})
	s.routes.Handle("/page/{page}", Handler{s, getIndex})
	ass := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/")))
	s.routes.PathPrefix("/assets/").Handler(ass)
	s.routes.Handle("/ero/{id}", Handler{s, getEro})
	s.routes.Handle("/login", Handler{s, getLogin})
	s.routes.Handle("/auth/login", Handler{s, postLogin}).Methods("POST")
	s.routes.Handle("/signup", Handler{s, getSignup})
	s.routes.Handle("/auth/signup", Handler{s, postSignup}).Methods("POST")

	s.routes.NotFoundHandler = Handler{s, getStop}

	// handle sigint
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case sig := <-c:
			fmt.Printf("\n[*] got %s signal. aborting...\n", sig)
			os.Exit(0)
		}
	}()
	log.Fatal(http.ListenAndServe(":1337", s.routes))
}
