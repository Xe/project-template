package main

import (
	"net/http"

	"git.xeserv.us/xe/project-template/models"
	"git.xeserv.us/xe/project-template/web"
	"git.xeserv.us/xe/project-template/xess"
	"github.com/a-h/templ"
)

type Server struct {
	dao *models.DAO
}

func (s *Server) register(mux *http.ServeMux) {
	xess.Mount(mux)
	mux.HandleFunc("/{$}", s.Index)
	mux.HandleFunc("/", s.NotFound)
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	templ.Handler(xess.Simple("Hello!", web.Index())).ServeHTTP(w, r)
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	xess.NotFound(w, r)
}
