package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"

	"git.xeserv.us/xe/project-template/web"
	"git.xeserv.us/xe/project-template/xess"
	"github.com/a-h/templ"
	"github.com/facebookgo/flagenv"
)

var (
	bind = flag.String("bind", ":4000", "TCP host:port to bind on")
)

func main() {
	flagenv.Parse()
	flag.Parse()

	mux := http.NewServeMux()

	xess.Mount(mux)

	mux.Handle("/{$}", templ.Handler(xess.Simple("Hello!", web.Index())))

	mux.HandleFunc("/", xess.NotFound)

	slog.Info("now listening", "url", "http://localhost"+*bind)
	log.Fatal(http.ListenAndServe(*bind, mux))
}
