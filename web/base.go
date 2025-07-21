package web

import (
	"embed"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/NYTimes/gziphandler"
	"github.com/Xe/project-template/internal"
)

var (
	//go:embed static
	static embed.FS
)

// isCompressible returns true if the content type should be compressed
func isCompressible(contentType string) bool {
	// Remove charset and other parameters
	mediaType := strings.Split(contentType, ";")[0]
	mediaType = strings.TrimSpace(mediaType)

	return strings.HasPrefix(mediaType, "text/") ||
		mediaType == "application/javascript" ||
		mediaType == "application/json" ||
		mediaType == "application/xml" ||
		mediaType == "application/rss+xml" ||
		mediaType == "application/atom+xml"
}

func selectiveGzipHandler(next http.Handler) http.Handler {
	gzipHandler := gziphandler.GzipHandler(next)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)
		contentType := mime.TypeByExtension(ext)

		if contentType != "" && isCompressible(contentType) {
			gzipHandler.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func Mount(mux *http.ServeMux) {
	var h http.Handler = http.FileServerFS(static)
	h = internal.UnchangingCache(h)
	h = selectiveGzipHandler(h)

	mux.Handle("/static/", h)
}
