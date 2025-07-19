package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"

	"git.xeserv.us/xe/project-template/internal"
	"github.com/facebookgo/flagenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	bind        = flag.String("bind", ":4000", "TCP host:port to bind on")
	databaseURL = flag.String("database-url", "", "Postgres database URL")
	redisURL    = flag.String("redis-url", "", "Valkey/Redis database URL")
	slogLevel   = flag.String("slog-level", "INFO", "log level")
)

func main() {
	flagenv.Parse()
	flag.Parse()

	internal.InitSlog(*slogLevel)

	s, err := New(Options{
		DatabaseURL: *databaseURL,
		RedisURL:    *redisURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	s.register(mux)

	var h http.Handler = mux
	h = s.sessionMiddleware(h)

	slog.Info("now listening", "url", "http://localhost"+*bind)
	log.Fatal(http.ListenAndServe(*bind, h))
}
