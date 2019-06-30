package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func postEdit(s *Server, w http.ResponseWriter, r *http.Request) error {
	var err error
	r.ParseForm()
	vars := mux.Vars(r)
	redir := fmt.Sprintf("/admin/edit/page/%s", r.PostFormValue("page"))
	if r.PostFormValue("group") == "new" {
		ng := r.PostFormValue("new-group")
		id := r.PostFormValue("id")
		if ng == "" {
			return StatusError{http.StatusTeapot,
				errors.New("type in the group idiot")}
		}
		_, err = s.DB.CreateMeta.Exec(ng, fmt.Sprintf(`["%s"]`, id), vars["id"])
		if err != nil {
			return err
		}
		http.Redirect(w, r, redir, http.StatusFound)
		return nil
	}
	g, i := r.PostFormValue("group"), r.PostFormValue("id")
	switch r.PostFormValue("op") {
	case "+":
		_, err := s.DB.UpdateMeta.Exec(
			fmt.Sprintf("{%s, -1}", g),
			fmt.Sprintf(`"%s"`, i),
			vars["id"],
		)
		if err != nil {
			return err
		}
	}
	http.Redirect(w, r, redir, http.StatusFound)
	return nil
}
