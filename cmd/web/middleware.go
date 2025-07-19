package main

import (
	"context"
	"net/http"

	"git.xeserv.us/xe/project-template/models"
	"github.com/gorilla/sessions"
)

type sessionKey struct{}

func getSession(ctx context.Context) (*sessions.Session, bool) {
	session, ok := ctx.Value(sessionKey{}).(*sessions.Session)
	return session, ok
}

func withSession(ctx context.Context, session *sessions.Session) context.Context {
	return context.WithValue(ctx, sessionKey{}, session)
}

type flashesKey struct{}

func getFlashes(ctx context.Context) ([]models.Flash, bool) {
	flashes, ok := ctx.Value(flashesKey{}).([]models.Flash)
	return flashes, ok
}

func withFlashes(ctx context.Context, sess *sessions.Session) context.Context {
	var flashes []models.Flash

	for _, thing := range sess.Flashes() {
		if flash, ok := thing.(models.Flash); ok {
			flashes = append(flashes, flash)
		}
	}

	return context.WithValue(ctx, flashesKey{}, flashes)
}

func (s *Server) sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Trailer", "Set-Cookie")

		r = r.WithContext(withSession(r.Context(), session))
		r = r.WithContext(withFlashes(r.Context(), session))

		next.ServeHTTP(w, r)

		if err := s.store.Save(r, w, session); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
