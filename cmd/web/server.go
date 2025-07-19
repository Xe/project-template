package main

import (
	"net/http"

	"git.xeserv.us/xe/project-template/models"
	"git.xeserv.us/xe/project-template/web"
	"github.com/a-h/templ"
	"github.com/rbcervilla/redisstore/v9"
)

const sessionName = "session"

type Server struct {
	dao   *models.DAO
	store *redisstore.RedisStore
}

func (s *Server) register(mux *http.ServeMux) {
	web.Mount(mux)
	mux.HandleFunc("/{$}", s.Index)
	mux.HandleFunc("/test/addflash/{kind}", s.AddFlash)
	mux.HandleFunc("/", s.NotFound)
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	flashes, _ := getFlashes(r.Context())

	templ.Handler(web.Simple("Hello!", web.Index(), flashes)).ServeHTTP(w, r)
}

func (s *Server) AddFlash(w http.ResponseWriter, r *http.Request) {
	session, _ := getSession(r.Context())

	var flash models.Flash

	switch r.PathValue("kind") {
	case string(models.FlashInfo):
		flash = models.NewFlash(models.FlashInfo, "<p>Welcome!</p>")
	case string(models.FlashFailure):
		flash = models.NewFlash(models.FlashFailure, "<p>Oh no, something went wrong!</p>")
	case string(models.FlashSuccess):
		flash = models.NewFlash(models.FlashSuccess, "<p>Oh yes, something went right!</p>")
	case string(models.FlashWarning):
		flash = models.NewFlash(models.FlashWarning, "<p>Are you sure you wanted to do that?</p>")
	default:
		templ.Handler(
			web.Simple("Not found: "+r.URL.Path, web.NotFound(r.URL.Path), nil),
			templ.WithStatus(http.StatusNotFound),
		).ServeHTTP(w, r)
		return
	}

	session.AddFlash(flash)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	templ.Handler(
		web.Simple("Not found: "+r.URL.Path, web.NotFound(r.URL.Path), nil),
		templ.WithStatus(http.StatusNotFound),
	).ServeHTTP(w, r)
}
