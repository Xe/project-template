package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Xe/project-template/models"
	"github.com/Xe/project-template/web"
	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v9"
)

const sessionName = "session"

type Options struct {
	DatabaseURL string
	RedisURL    string
}

func New(opts Options) (*Server, error) {
	rdb, err := models.ConnectValkey(opts.RedisURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to valkey: %w", err)
	}

	dao, err := models.New(opts.DatabaseURL, rdb)
	if err != nil {
		return nil, fmt.Errorf("can't create DAO: %w", err)
	}

	store, err := redisstore.NewRedisStore(context.Background(), rdb)
	if err != nil {
		return nil, fmt.Errorf("can't create redis store: %w", err)
	}

	store.KeyPrefix("session:")
	store.Options(sessions.Options{
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   86400 * 60,
	})

	result := &Server{
		dao:   dao,
		store: store,
	}

	return result, nil
}

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
