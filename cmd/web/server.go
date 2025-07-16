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
	mux.HandleFunc("/", s.NotFound)
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	session, err := s.store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.store.Save(r, w, session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ.Handler(web.Simple("Hello!", web.Index())).ServeHTTP(w, r)
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	templ.Handler(
		web.Simple("Not found: "+r.URL.Path, web.NotFound(r.URL.Path)),
		templ.WithStatus(http.StatusNotFound),
	).ServeHTTP(w, r)
}
