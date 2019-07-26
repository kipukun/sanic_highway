package http

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
	"github.com/kipukun/sanic_highway/model"
)

func postEdit(s *Server, w http.ResponseWriter, r *http.Request) error {
	_, err := authenticate(s, w, r)
	if err != nil {
		return nil
	}
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
	case "-":
		_, err := s.DB.RemoveMeta.Exec(
			fmt.Sprintf("{%s}", g),
			fmt.Sprintf("%s", g),
			vars["id"],
		)
		if err != nil {
			return err
		}
	case "delete":
		_, err := s.DB.DeleteMeta.Exec(
			fmt.Sprintf("%s", g),
			vars["id"],
		)
		if err != nil {
			return err
		}
	}
	http.Redirect(w, r, redir, http.StatusFound)
	return nil
}

func postIngest(s *Server, w http.ResponseWriter, r *http.Request) error {
	_, err := authenticate(s, w, r)
	if err != nil {
		return nil
	}
	var buf bytes.Buffer
	f, _, err := r.FormFile("file")
	if err != nil {
		return err
	}
	io.Copy(&buf, f)

	// we don't need the contents anymore
	f.Close()
	err = s.DB.Ingest(&buf)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/admin/edit", http.StatusFound)
	return nil
}

func postExport(s *Server, w http.ResponseWriter, r *http.Request) error {
	_, err := authenticate(s, w, r)
	if err != nil {
		return nil
	}
	var buf bytes.Buffer
	z := gzip.NewWriter(&buf)
	var ero []model.Eroge
	err = s.DB.All.Select(&ero)
	if err != nil {
		return err
	}
	b, err := json.Marshal(ero)
	if err != nil {
		return err
	}
	_, err = z.Write(b)
	if err != nil {
		return err
	}
	z.Close()
	w.Header().Set("Content-Disposition", "attachment; filename=export.json.gz")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
	w.Write(buf.Bytes())
	return nil
}
