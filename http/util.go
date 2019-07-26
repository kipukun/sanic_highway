package http

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kipukun/sanic_highway/model"
)

// authenticate takes in a Request and returns a valid User and nil.
// otherwise, returns an empty object and the error.
func authenticate(s *Server, w http.ResponseWriter, r *http.Request) (*model.User, error) {
	u := &model.User{}
	c, err := r.Cookie("id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return u, err
	}
	err = s.DB.Lookup.Get(u, c.Value)
	if err != nil {
		return u, err
	}
	return u, nil
}

// generates page numbers for a given "/page/N" variable with a Request.
func paginate(r *http.Request) (int, []int, error) {
	var pg int
	var err error
	pn := make([]int, 5)
	vars := mux.Vars(r)

	if vars["page"] == "" {
		pg = 0
	} else {
		pg, err = strconv.Atoi(vars["page"])
		if err != nil {
			return 0, pn, err
		}
	}

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
	return pg, pn, nil
}
