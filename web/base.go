package web

import (
	"embed"
	"net/http"

	"git.xeserv.us/xe/project-template/internal"
)

var (
	//go:embed static
	static embed.FS
)

func Mount(mux *http.ServeMux) {
	mux.Handle("/static/", internal.UnchangingCache(http.FileServerFS(static)))
}
