package internal

import (
	"net/http"

	"git.xeserv.us/xe/project-template/globals"
)

func UnchangingCache(h http.Handler) http.Handler {
	if globals.Version == "devel" {
		return h
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		h.ServeHTTP(w, r)
	})
}
