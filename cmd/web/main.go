package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"

	"git.xeserv.us/xe/project-template/internal"
	"git.xeserv.us/xe/project-template/models"
	"git.xeserv.us/xe/project-template/web"
	"git.xeserv.us/xe/project-template/xess"
	"github.com/a-h/templ"
	"github.com/facebookgo/flagenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	bind        = flag.String("bind", ":4000", "TCP host:port to bind on")
	databaseURL = flag.String("database-url", "", "Postgres database URL")
	slogLevel   = flag.String("slog-level", "INFO", "log level")
)

func main() {
	flagenv.Parse()
	flag.Parse()

	internal.InitSlog(*slogLevel)

	dao, err := models.New(*databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	_ = dao

	mux := http.NewServeMux()

	xess.Mount(mux)

	mux.Handle("/{$}", templ.Handler(xess.Simple("Hello!", web.Index())))

	mux.HandleFunc("/", xess.NotFound)

	slog.Info("now listening", "url", "http://localhost"+*bind)
	log.Fatal(http.ListenAndServe(*bind, mux))
}
