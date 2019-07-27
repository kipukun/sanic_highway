package http

import (
	"fmt"
	"net/http"
	"os/signal"

	"time"

	"os"

	"github.com/gorilla/mux"
	"github.com/kipukun/sanic_highway/config"
	"github.com/kipukun/sanic_highway/db"
)

// Server holds the DB connection, and configured router for the http package.
type Server struct {
	DB     *db.Database
	router *mux.Router
	HS     *http.Server
	Config config.Config
}

// Init takes in a Database and returns a new Server intialized with the
// default mux.Router and the database connection.
func Init(d *db.Database, c config.Config) *Server {
	r := mux.NewRouter()
	h := &http.Server{
		Handler:      r,
		Addr:         c.Web.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv := &Server{
		DB:     d,
		router: r,
		HS:     h,
		Config: c,
	}
	return srv
}

// Start sets up router on the mux and starts the HTTP server.
func (s *Server) Start() {
	// basic routes
	s.router.Handle("/", Handler{s, getIndex})
	s.router.Handle("/about", Handler{s, getAbout})
	s.router.Handle("/page/{page}", Handler{s, getIndex})
	ass := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/")))
	s.router.PathPrefix("/assets/").Handler(ass)
	s.router.Handle("/ero/{id}", Handler{s, getEro})
	// admin routes
	s.router.Handle("/admin", Handler{s, getAdmin})
	s.router.Handle("/admin/edit", Handler{s, getAdminEdit})
	s.router.Handle("/admin/edit/page/{page}", Handler{s, getAdminEdit})
	// auth routes
	s.router.Handle("/login", Handler{s, getLogin})
	s.router.Handle("/auth/login", Handler{s, postLogin}).Methods("POST")

	// api routes
	s.router.Handle("/api/edit/{id}", Handler{s, postEdit}).Methods("POST")
	s.router.Handle("/api/ingest", Handler{s, postIngest}).Methods("POST")
	s.router.Handle("/api/export", Handler{s, postExport}).Methods("POST")

	s.router.NotFoundHandler = Handler{s, getStop}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case sig := <-c:
			fmt.Printf("\n[*] got %s signal. shutting down...\n", sig)
			s.HS.Close()
		}
	}()
	err := s.HS.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("could not close gracefully")
	}
}
