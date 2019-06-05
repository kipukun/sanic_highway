package http

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kipukun/sanic_highway/model"
	"github.com/kipukun/sanic_highway/templates"
	"golang.org/x/crypto/bcrypt"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Handler struct {
	*Server
	H func(s *Server, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP satisifies the http.Handler interface.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Server, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func getIndex(s *Server, w http.ResponseWriter, r *http.Request) error {
	var erogeList []model.Eroge
	var pg int
	var err error
	var pn [5]int
	index := true

	if index {
		pg = 0
	} else {
		vars := mux.Vars(r)
		pg, err = strconv.Atoi(vars["page"])
		if err != nil {
			return StatusError{http.StatusInternalServerError, err}
		}
	}

	err = s.DB.Eros.Select(&erogeList, pg*50)
	if err != nil {
		return StatusError{http.StatusInternalServerError, err}
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

	return nil
}

func getAbout(s *Server, w http.ResponseWriter, r *http.Request) error {
	p := &templates.AboutPage{}
	templates.WritePageTemplate(w, p)
	return nil
}

func getStop(s *Server, w http.ResponseWriter, r *http.Request) error {
	p := &templates.StopPage{}
	templates.WritePageTemplate(w, p)
	return nil
}

func getEro(s *Server, w http.ResponseWriter, r *http.Request) error {
	var eroge model.Eroge
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return StatusError{http.StatusInternalServerError, err}

	}
	err = s.DB.Ero.Get(&eroge, id)
	if err != nil {
		p := &templates.StopPage{}
		templates.WritePageTemplate(w, p)
		return StatusError{http.StatusNotFound, err}
	}
	p := &templates.ErogePage{
		Ero: eroge,
	}

	templates.WritePageTemplate(w, p)

	return nil
}
func getLogin(s *Server, w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte(`
		<html><head><title>sekrit</title></head>
		<body>
		<form action="/auth/login" method="post">
			<input type="text" name="user" value="user" required><br />
			<input type="password" name="pass" value="pass" required><br />
			<input type="submit" value="submit">
		</form>
		</body></html>
		`))
	return nil
}
func postLogin(s *Server, w http.ResponseWriter, r *http.Request) error {
	return nil
}
func getSignup(s *Server, w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte(`
		<html><head><title>sekrit</title></head>
		<body>
		<form action="/auth/signup" method="post">
			<input type="text" name="user" value="user" required><br />
			<input type="password" name="pass" placeholder="pass" required><br />
			<input type="password" name="conf" placeholder="pass" required><br />
			<input type="submit" value="submit">
		</form>
		</body></html>
		`))
	return nil
}
func postSignup(s *Server, w http.ResponseWriter, r *http.Request) error {
	user, pass := r.PostFormValue("user"), r.PostFormValue("pass")
	if pass != r.PostFormValue("conf") {
		return StatusError{http.StatusNotAcceptable,
			errors.New(`Password does not match confirmation.`)}
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return StatusError{http.StatusInternalServerError, err}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		return StatusError{http.StatusInternalServerError, err}
	}
	_, err = s.DB.InsertUser.Exec(id, user, string(hash))
	if err != nil {
		return StatusError{http.StatusInternalServerError, err}
	}
	return nil
}
