// Package xess vendors a copy of Xess and makes it available at /static/xess/xess.css
//
// This is intended to be used as a vendored package in other projects.
package xess

import (
	"embed"
	"net/http"

	"git.xeserv.us/xe/project-template/globals"
	"git.xeserv.us/xe/project-template/internal"
	"github.com/a-h/templ"
)

//go:generate go tool templ generate

var (
	//go:embed xess.css static
	Static embed.FS

	URL = "/static/xess/xess.css"
)

func init() {
	Mount(http.DefaultServeMux)

	URL = URL + "?cachebuster=" + globals.Version
}

func Mount(mux *http.ServeMux) {
	mux.Handle("/static/xess/", internal.UnchangingCache(http.StripPrefix("/static/xess/", http.FileServerFS(Static))))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	templ.Handler(
		Simple("Not found: "+r.URL.Path, fourohfour(r.URL.Path)),
		templ.WithStatus(http.StatusNotFound),
	)
}
